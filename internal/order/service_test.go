package order_test

import (
	"math"
	"testing"

	"github.com/kveriz/algo/14/internal/order"
)

func TestOrderService_CalculateTotal(t *testing.T) {
	tests := []struct {
		name       string
		items      []order.Item
		clientType order.ClientType
		want       float64
	}{
		{
			name: "regular client, no threshold",
			items: []order.Item{
				order.NewItem("Book", 100, 2),
			},
			clientType: order.ClientTypeRegular,
			want:       200,
		},
		{
			name: "VIP discount",
			items: []order.Item{
				order.NewItem("Book", 100, 2),
			},
			clientType: order.ClientTypeVIP,
			want:       180,
		},
		{
			name: "NEW discount",
			items: []order.Item{
				order.NewItem("Book", 100, 2),
			},
			clientType: order.ClientTypeNew,
			want:       190,
		},
		{
			name: "threshold after regular sum",
			items: []order.Item{
				order.NewItem("Phone", 600, 2),
			},
			clientType: order.ClientTypeRegular,
			want:       1150,
		},
		{
			name: "VIP then threshold",
			items: []order.Item{
				order.NewItem("Phone", 600, 2),
			},
			clientType: order.ClientTypeVIP,
			want:       1030,
		},
		{
			name: "NEW then threshold",
			items: []order.Item{
				order.NewItem("Phone", 600, 2),
			},
			clientType: order.ClientTypeNew,
			want:       1090,
		},
		{
			name: "VIP discount prevents threshold",
			items: []order.Item{
				order.NewItem("Laptop", 1050, 1),
			},
			clientType: order.ClientTypeVIP,
			want:       945,
		},
		{
			name: "NEW discount then threshold",
			items: []order.Item{
				order.NewItem("Laptop", 1100, 1),
			},
			clientType: order.ClientTypeNew,
			want:       995,
		},
		// DTT
		{
			name: "exactly 1000 does not apply threshold",
			items: []order.Item{
				order.NewItem("Laptop", 1000, 1),
			},
			clientType: order.ClientTypeRegular,
			want:       1000,
		},
		{
			name:       "empty items",
			items:      []order.Item{},
			clientType: order.ClientTypeVIP,
			want:       0,
		},
		{
			name: "multiple items",
			items: []order.Item{
				order.NewItem("Book", 100, 1),
				order.NewItem("Pen", 200, 2),
			},
			clientType: order.ClientTypeVIP,
			want:       450,
		},
		{
			name: "unknown client type",
			items: []order.Item{
				order.NewItem("Book", 100, 2),
			},
			clientType: order.ClientType("GOLD"),
			want:       200,
		},
		{
			name: "exactly 10 units — no bulk discount",
			items: []order.Item{
				order.NewItem("Book", 100, 10),
			},
			clientType: order.ClientTypeRegular,
			want:       1000,
		},
		{
			name: "11 units — bulk discount applied",
			items: []order.Item{
				order.NewItem("Book", 100, 11),
			},
			clientType: order.ClientTypeRegular,
			want:       1039, // 1100 * 0.99 = 1089 → −50 = 1039
		},
		{
			name: "bulk discount prevents threshold",
			items: []order.Item{
				order.NewItem("Laptop", 1000, 1),
				order.NewItem("Pen", 0.5, 10),
			},
			clientType: order.ClientTypeRegular,
			want:       994.95, // 1005 * 0.99 = 994.95 → порог не срабатывает
		},
		{
			name: "bulk discount combined with VIP",
			items: []order.Item{
				order.NewItem("Book", 100, 11),
			},
			clientType: order.ClientTypeVIP,
			want:       980.1, // 1100 * 0.9 = 990 → * 0.99 = 980.1
		},
		{
			name: "bulk discount combined with NEW",
			items: []order.Item{
				order.NewItem("Book", 100, 11),
			},
			clientType: order.ClientTypeNew,
			want:       984.55, // 1100 * 0.95 = 1045 → * 0.99 = 1034.55 → −50 = 984.55
		},
	}

	service := order.NewOrderService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.CalculateTotal(tt.items, tt.clientType)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Fatalf("CalculateTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
