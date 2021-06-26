package kafkaproducer

import (
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/Shopify/sarama"
	"github.com/urfave/cli"
	"google.golang.org/protobuf/proto"

	votingpb "github.com/egsam98/voting/proto"
)

const (
	structVote  = "Vote"
	structVoter = "Voter"
)

var Cmd = cli.Command{
	Name:  "kafka-producer",
	Usage: "send proto message to Kafka using JSON input",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "struct-name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "json",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "kafka-addr",
			Value: "localhost:29092",
		},
		&cli.StringFlag{
			Name:     "kafka-topic",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		return run(
			ctx.String("struct-name"),
			[]byte(ctx.String("json")),
			ctx.String("kafka-addr"),
			ctx.String("kafka-topic"),
		)
	},
}

func run(structName string, jsonStr []byte, kafkaAddr, topic string) error {
	var protoMsg proto.Message
	switch structName {
	case structVote:
		protoMsg = &votingpb.Vote{}
	case structVoter:
		protoMsg = &votingpb.Voter{}
	default:
		return errors.New("unknown struct name")
	}

	if err := protojson.Unmarshal(jsonStr, protoMsg); err != nil {
		return errors.Wrapf(err, "failed to unmarshal JSON to %T", protoMsg)
	}

	b, err := proto.Marshal(protoMsg)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal %T to protobuf", protoMsg)
	}

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{kafkaAddr}, cfg)
	if err != nil {
		return errors.Wrap(err, "failed to init Kafka producer")
	}

	if _, _, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(b),
	}); err != nil {
		return errors.Wrapf(err, "failed to send protobuf message to topic %q", topic)
	}

	fmt.Printf("message has been sent successfully to topic %q\n", topic)
	return nil
}
