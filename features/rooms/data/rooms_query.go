package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"project/features/instruments"
	"project/features/rooms"
	"project/helper"

	"firebase.google.com/go/db"
	"github.com/sashabaranov/go-openai"
)

type RoomData struct {
	db     *db.Client
	openai *openai.Client
}

func NewRoomData(client *db.Client, openai *openai.Client) rooms.RoomDataInterface {
	return &RoomData{
		db:     client,
		openai: openai,
	}
}

func (rd *RoomData) GetAllRooms() (map[string]map[string]interface{}, error) {
	ref := rd.db.NewRef("rooms")
	rooms := map[string]map[string]interface{}{}
	if err := ref.Get(context.Background(), &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rd *RoomData) GetRoomByID(roomID string) (*rooms.Rooms, error) {
	ref := rd.db.NewRef("rooms").Child(roomID)
	var room rooms.Rooms
	if err := ref.Get(context.Background(), &room); err != nil {
		return nil, err
	}
	return &room, nil
}

func (rd *RoomData) AddRoom(newRoom rooms.Rooms) (*rooms.Rooms, error) {
	if roomExists := rd.IsRoomExist(newRoom.RoomID); roomExists {
		return nil, errors.New("room already exists")
	}
	ref := rd.db.NewRef("rooms").Child(newRoom.RoomID)
	if err := ref.Set(context.Background(), newRoom); err != nil {
		return nil, err
	}
	return &newRoom, nil
}

func (rd *RoomData) DeleteRoom(roomID string) error {
	ref := rd.db.NewRef("rooms").Child(roomID)
	if err := ref.Delete(context.Background()); err != nil {
		return err
	}
	return nil
}

func (rd *RoomData) UpdateRoom(roomID string, updatedRoom rooms.Rooms) (*rooms.Rooms, error) {
	ref := rd.db.NewRef("rooms").Child(roomID)
	res, err := helper.ToMap(updatedRoom)
	if err != nil {
		return nil, err
	}

	if err := ref.Update(context.Background(), res); err != nil {
		return nil, err
	}
	return &updatedRoom, nil
}

func (rd *RoomData) FilterRoomByPrice(price int) (map[string]map[string]interface{}, error) {
	rooms := map[string]map[string]interface{}{}
	ref := rd.db.NewRef("rooms")
	if err := ref.OrderByChild("price").EndAt(price).Get(context.Background(), &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rd *RoomData) GetBookedRooms() (map[string]map[string]interface{}, error) {
	reserved := map[string]map[string]interface{}{}
	ref := rd.db.NewRef("reservations")
	err := ref.Get(context.Background(), &reserved)
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
		return nil, err
	}
	return reserved, nil
}

func (rd *RoomData) IsRoomExist(roomID string) bool {
	ref := rd.db.NewRef("rooms").Child(roomID)
	var room map[string]any
	ref.Get(context.Background(), &room)
	fmt.Println("room : ", room)
	if room == nil {
		return false
	}
	return true
}

func (rd *RoomData) AddRoomInstrument(roomID string, instrumentData instruments.RoomInstrument) (any, error) {
	ref := rd.db.NewRef("rooms/" + roomID).Child("instrument")
	instrument := make(map[string]any)

	bassExist := rd.isInstrumentExists(instrumentData.Bass)
	if bassExist {
		instrument["bass"] = instrumentData.Bass
	}

	guitarExist := rd.isInstrumentExists(instrumentData.Guitar)
	if guitarExist {
		instrument["guitar"] = instrumentData.Guitar
	}

	keyboardExist := rd.isInstrumentExists(instrumentData.Keyboard)
	if keyboardExist {
		instrument["keyboard"] = instrumentData.Keyboard
	}

	drumExist := rd.isInstrumentExists(instrumentData.Drum)
	if drumExist {
		instrument["drum"] = instrumentData.Drum
	}

	if !bassExist && !guitarExist && !keyboardExist && !drumExist {
		return nil, errors.New("no instruments found")
	}
	if err := ref.Update(context.Background(), instrument); err != nil {
		return nil, err
	}
	return instrument, nil
}

func (rd *RoomData) isInstrumentExists(instrumentName string) bool {
	if len(instrumentName) < 1 {
		return false
	}
	ref := rd.db.NewRef("instruments").Child(instrumentName)
	var instrument map[string]any
	ref.Get(context.Background(), &instrument)
	if instrument == nil {
		return false
	}
	return true
}

func (rd *RoomData) GetRecommendation(genre1 string, genre2 string, message string) (string, error) {
	str := "\\ruangB\\"
	res, err := rd.openai.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: "you're here to help people choose their rooms from your recommendations based on data given. for the room names, don't include the / into your response, don't need to make a new line, don't need to include double escape or if you find something like this : " + str + ", remove the backslash, just straight up room names and its instruments like 'ruangH - fender stratocaster. reasons....' in a single paragraph"},
			{Role: openai.ChatMessageRoleUser, Content: message},
		},
	})

	if err != nil {
		fmt.Println("data error")
		return "", fmt.Errorf("ChatCompletion error: %v\n", err)
	}

	return res.Choices[0].Message.Content, nil
}
