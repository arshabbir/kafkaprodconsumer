package clients

import (
	"context"
	"reflect"
	"testing"
)

func TestNewProducer(t *testing.T) {
	type args struct {
		brokers []string
		topic   string
	}
	tests := []struct {
		name string
		args args
		want KafkaProducer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProducer(tt.args.brokers, tt.args.topic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProducer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kafkaProducer_SendMessage(t *testing.T) {
	type args struct {
		ctx context.Context
		msg []byte
	}
	tests := []struct {
		name    string
		p       *kafkaProducer
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.SendMessage(tt.args.ctx, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("kafkaProducer.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
