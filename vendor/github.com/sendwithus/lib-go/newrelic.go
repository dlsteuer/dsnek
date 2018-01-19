package swu

import (
	"context"
	"github.com/newrelic/go-agent"
)

type TracingAgent interface {
	StartTransaction(txnName string) Transaction
	StartSegment(segmentName string, txn Transaction) Segment
}

type Transaction interface {
	StartSegment(segmentName string) Segment
	End()
}

type Segment interface {
	End()
}

type noopAgent struct{}
type noopTransaction struct{}
type noopSegment struct{}

func (a *noopAgent) StartTransaction(txnName string) Transaction {
	return &noopTransaction{}
}

func (a *noopAgent) StartSegment(segmentName string, txn Transaction) Segment {
	return &noopSegment{}
}

func (t *noopTransaction) StartSegment(segmentName string) Segment {
	return &noopSegment{}
}

func (t *noopTransaction) End() {}
func (s *noopSegment) End()     {}

type newRelicAgent struct {
	newRelicApp newrelic.Application
}

func (a *newRelicAgent) StartTransaction(txnName string) Transaction {
	txn := a.newRelicApp.StartTransaction(txnName, nil, nil)
	return &newRelicTransaction{newRelicTransaction: txn}
}

func (a *newRelicAgent) StartSegment(segmentName string, txn Transaction) Segment {
	seg := newrelic.StartSegment((txn.(*newRelicTransaction)).newRelicTransaction, segmentName)
	return &newRelicSegment{newRelicSegment: seg}
}

type newRelicTransaction struct {
	newRelicTransaction newrelic.Transaction
}

func (t *newRelicTransaction) StartSegment(segmentName string) Segment {
	seg := newrelic.StartSegment(t.newRelicTransaction, segmentName)
	return &newRelicSegment{newRelicSegment: seg}
}

func (t *newRelicTransaction) End() {
	t.newRelicTransaction.End()
}

type newRelicSegment struct {
	newRelicSegment newrelic.Segment
}

func (t *newRelicSegment) End() {
	t.newRelicSegment.End()
}

func NewNewRelicAgent(appName, licenseKey *string) TracingAgent {
	if licenseKey == nil {
		return &noopAgent{}
	}

	config := newrelic.NewConfig(*appName, *licenseKey)
	app, err := newrelic.NewApplication(config)
	if err != nil {
		internalLogger.ErrorWithError(err, "Unable to initialize newrelic agent")
		return &noopAgent{}
	}
	return &newRelicAgent{newRelicApp: app}
}

func ContextWithTransaction(ctx context.Context, tx Transaction) context.Context {
	return context.WithValue(ctx, "tracingAgent", tx)
}

func ContextGetTransaction(ctx context.Context) Transaction {
	if val, ok := ctx.Value("tracingAgent").(Transaction); ok {
		return val
	}
	return &noopTransaction{}
}
