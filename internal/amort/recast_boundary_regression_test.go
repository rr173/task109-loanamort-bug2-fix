package amort_test

import (
	"testing"

	"task109-loanamort/internal/amort"
	"task109-loanamort/internal/domain"
)

func TestReducePaymentUsesRemainingTermAfterPaidPeriods(t *testing.T) {
	res, err := amort.RecastReducePayment(amort.RecastInput{Type: domain.EqualInstallment, Outstanding: 90000, PeriodicRateMicro: 5000, PaidPeriods: 1, OriginalTerm: 12})
	if err != nil {
		t.Fatal(err)
	}
	if res.TermPeriods != 12 {
		t.Fatalf("term=%d, want 12", res.TermPeriods)
	}
	if len(res.Schedule) != 11 {
		t.Fatalf("remaining periods=%d, want 11", len(res.Schedule))
	}
}
