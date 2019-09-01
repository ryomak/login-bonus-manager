package handler

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ryomak/login-bonus-manager/line-bot/src/line"
)

func LineHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	myLineRequest, err := line.UnmarshalLineRequest([]byte(request.Body))
	if err != nil {
		log.Fatal(err)
	}

	bot, err := linebot.New(
		os.Getenv("channelSecret"),
		os.Getenv("channelToken"),
	)
	var tmpReplyMessage string
	tmpReplyMessage = "回答：" + myLineRequest.Events[0].Message.Text
	if _, err = bot.ReplyMessage(myLineRequest.Events[0].ReplyToken, linebot.NewTextMessage(tmpReplyMessage)).Do(); err != nil {
		log.Fatal(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       request.Body,
	}, nil
}