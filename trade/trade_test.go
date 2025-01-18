package trade

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/billfort/binance-usdmfuture/pub"
)

// go test -v -run TestTestOrder
func TestTestOrder(t *testing.T) {
	type args struct {
		key *pub.Key
		op  *OrderParam
	}
	tests := []struct {
		name    string
		args    *args
		want    *orderResponse
		wantErr bool
	}{
		{
			"test",
			&args{
				key: &pub.Key{
					ApiKey:    "test",
					SecretKey: "test",
				},
				op: &OrderParam{
					Symbol:           "test",
					Side:             "test",
					Type:             "test",
					Quantity:         "0",
					Price:            "0",
					StopPrice:        "0",
					TimeInForce:      pub.TIF_GTC,
					NewClientOrderId: "test",
					NewOrderRespType: pub.RT_Ack,
					RecvWindow:       0,
					Timestamp:        0,
				},
			},
			&orderResponse{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TestOrder(tt.args.key, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -run TestNewOrder
func TestNewOrder(t *testing.T) {
	type args struct {
		key *pub.Key
		op  *OrderParam
	}
	tests := []struct {
		name    string
		args    *args
		wantErr error
	}{
		{
			"test",
			&args{
				key: pub.TestKey,
				op: &OrderParam{
					Symbol:           "BTCUSDT",
					Side:             pub.OS_Buy,
					PositionSide:     pub.PS_Long,
					Type:             pub.OT_Market,
					Quantity:         "0.01",
					NewClientOrderId: fmt.Sprintf("%d", time.Now().UnixMilli()),
				},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrder(tt.args.key, tt.args.op)
			if nil != tt.wantErr {
				t.Errorf("NewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("NewOrder() Response: %+v", got)
		})
	}
}

// go test -v -run TestBatchOrders
func TestBatchOrders(t *testing.T) {
	type args struct {
		key *pub.Key
		ops []OrderParam
	}
	tests := []struct {
		name    string
		args    *args
		wantErr error
	}{
		{
			"test",
			&args{
				key: pub.TestKey,
				ops: []OrderParam{
					{
						Symbol:           "BTCUSDT",
						Side:             pub.OS_Buy,
						PositionSide:     pub.PS_Long,
						Type:             pub.OT_Market,
						Quantity:         "0.01",
						NewClientOrderId: fmt.Sprintf("%d", time.Now().UnixMilli()),
					},
					{
						Symbol:           "BTCUSDT",
						Side:             pub.OS_Buy,
						PositionSide:     pub.PS_Long,
						Type:             pub.OT_Market,
						Quantity:         "0.01",
						NewClientOrderId: fmt.Sprintf("%d", time.Now().UnixMilli()),
					},
				},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BatchOrders(tt.args.key, tt.args.ops)
			if nil != tt.wantErr {
				t.Errorf("BatchOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("BatchOrders() Response: %+v", got)
		})
	}
}
