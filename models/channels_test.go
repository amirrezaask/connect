// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testChannels(t *testing.T) {
	t.Parallel()

	query := Channels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testChannelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChannelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Channels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChannelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChannelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChannelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ChannelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Channel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ChannelExists to return true, but got false.")
	}
}

func testChannelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	channelFound, err := FindChannel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if channelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testChannelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Channels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testChannelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Channels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testChannelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	channelOne := &Channel{}
	channelTwo := &Channel{}
	if err = randomize.Struct(seed, channelOne, channelDBTypes, false, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}
	if err = randomize.Struct(seed, channelTwo, channelDBTypes, false, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = channelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = channelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Channels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testChannelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	channelOne := &Channel{}
	channelTwo := &Channel{}
	if err = randomize.Struct(seed, channelOne, channelDBTypes, false, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}
	if err = randomize.Struct(seed, channelTwo, channelDBTypes, false, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = channelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = channelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func channelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Channel) error {
	*o = Channel{}
	return nil
}

func channelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Channel) error {
	*o = Channel{}
	return nil
}

func channelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Channel) error {
	*o = Channel{}
	return nil
}

func channelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Channel) error {
	*o = Channel{}
	return nil
}

func channelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Channel) error {
	*o = Channel{}
	return nil
}

func channelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Channel) error {
	*o = Channel{}
	return nil
}

func channelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Channel) error {
	*o = Channel{}
	return nil
}

func channelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Channel) error {
	*o = Channel{}
	return nil
}

func channelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Channel) error {
	*o = Channel{}
	return nil
}

func testChannelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Channel{}
	o := &Channel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, channelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Channel object: %s", err)
	}

	AddChannelHook(boil.BeforeInsertHook, channelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	channelBeforeInsertHooks = []ChannelHook{}

	AddChannelHook(boil.AfterInsertHook, channelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	channelAfterInsertHooks = []ChannelHook{}

	AddChannelHook(boil.AfterSelectHook, channelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	channelAfterSelectHooks = []ChannelHook{}

	AddChannelHook(boil.BeforeUpdateHook, channelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	channelBeforeUpdateHooks = []ChannelHook{}

	AddChannelHook(boil.AfterUpdateHook, channelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	channelAfterUpdateHooks = []ChannelHook{}

	AddChannelHook(boil.BeforeDeleteHook, channelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	channelBeforeDeleteHooks = []ChannelHook{}

	AddChannelHook(boil.AfterDeleteHook, channelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	channelAfterDeleteHooks = []ChannelHook{}

	AddChannelHook(boil.BeforeUpsertHook, channelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	channelBeforeUpsertHooks = []ChannelHook{}

	AddChannelHook(boil.AfterUpsertHook, channelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	channelAfterUpsertHooks = []ChannelHook{}
}

func testChannelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChannelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(channelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChannelToManyChannelPermissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Channel
	var b, c ChannelPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, channelPermissionDBTypes, false, channelPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, channelPermissionDBTypes, false, channelPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ChannelID = a.ID
	c.ChannelID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ChannelPermissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ChannelID == b.ChannelID {
			bFound = true
		}
		if v.ChannelID == c.ChannelID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ChannelSlice{&a}
	if err = a.L.LoadChannelPermissions(ctx, tx, false, (*[]*Channel)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChannelPermissions = nil
	if err = a.L.LoadChannelPermissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testChannelToManyUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Channel
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"channel_users\" (\"channel_id\", \"user_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"channel_users\" (\"channel_id\", \"user_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Users().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ChannelSlice{&a}
	if err = a.L.LoadUsers(ctx, tx, false, (*[]*Channel)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Users = nil
	if err = a.L.LoadUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testChannelToManyMessages(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Channel
	var b, c Message

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ChannelID = a.ID
	c.ChannelID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Messages().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ChannelID == b.ChannelID {
			bFound = true
		}
		if v.ChannelID == c.ChannelID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ChannelSlice{&a}
	if err = a.L.LoadMessages(ctx, tx, false, (*[]*Channel)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Messages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Messages = nil
	if err = a.L.LoadMessages(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Messages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testChannelToManyAddOpChannelPermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Channel
	var b, c, d, e ChannelPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, channelDBTypes, false, strmangle.SetComplement(channelPrimaryKeyColumns, channelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ChannelPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, channelPermissionDBTypes, false, strmangle.SetComplement(channelPermissionPrimaryKeyColumns, channelPermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ChannelPermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChannelPermissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChannelID {
			t.Error("foreign key was wrong value", a.ID, first.ChannelID)
		}
		if a.ID != second.ChannelID {
			t.Error("foreign key was wrong value", a.ID, second.ChannelID)
		}

		if first.R.Channel != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Channel != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ChannelPermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChannelPermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChannelPermissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testChannelToManyAddOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Channel
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, channelDBTypes, false, strmangle.SetComplement(channelPrimaryKeyColumns, channelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Channels[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Channels[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Users[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Users[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Users().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testChannelToManySetOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Channel
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, channelDBTypes, false, strmangle.SetComplement(channelPrimaryKeyColumns, channelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetUsers(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUsers(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Channels) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Channels) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Channels[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Channels[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Users[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Users[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testChannelToManyRemoveOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Channel
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, channelDBTypes, false, strmangle.SetComplement(channelPrimaryKeyColumns, channelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUsers(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUsers(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Channels) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Channels) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Channels[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Channels[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Users) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Users[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Users[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testChannelToManyAddOpMessages(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Channel
	var b, c, d, e Message

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, channelDBTypes, false, strmangle.SetComplement(channelPrimaryKeyColumns, channelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Message{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Message{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMessages(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChannelID {
			t.Error("foreign key was wrong value", a.ID, first.ChannelID)
		}
		if a.ID != second.ChannelID {
			t.Error("foreign key was wrong value", a.ID, second.ChannelID)
		}

		if first.R.Channel != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Channel != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Messages[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Messages[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Messages().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testChannelToOneHubUsingHub(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Channel
	var foreign Hub

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, channelDBTypes, false, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, hubDBTypes, false, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.HubID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Hub().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ChannelSlice{&local}
	if err = local.L.LoadHub(ctx, tx, false, (*[]*Channel)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Hub == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Hub = nil
	if err = local.L.LoadHub(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Hub == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testChannelToOneSetOpHubUsingHub(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Channel
	var b, c Hub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, channelDBTypes, false, strmangle.SetComplement(channelPrimaryKeyColumns, channelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, hubDBTypes, false, strmangle.SetComplement(hubPrimaryKeyColumns, hubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, hubDBTypes, false, strmangle.SetComplement(hubPrimaryKeyColumns, hubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Hub{&b, &c} {
		err = a.SetHub(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Hub != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Channels[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.HubID != x.ID {
			t.Error("foreign key was wrong value", a.HubID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.HubID))
		reflect.Indirect(reflect.ValueOf(&a.HubID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.HubID != x.ID {
			t.Error("foreign key was wrong value", a.HubID, x.ID)
		}
	}
}

func testChannelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testChannelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChannelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testChannelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Channels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	channelDBTypes = map[string]string{`ID`: `character varying`, `Name`: `character varying`, `Type`: `enum.channel_type_t('voice','text')`, `HubID`: `character varying`}
	_              = bytes.MinRead
)

func testChannelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(channelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(channelAllColumns) == len(channelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, channelDBTypes, true, channelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testChannelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(channelAllColumns) == len(channelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Channel{}
	if err = randomize.Struct(seed, o, channelDBTypes, true, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, channelDBTypes, true, channelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(channelAllColumns, channelPrimaryKeyColumns) {
		fields = channelAllColumns
	} else {
		fields = strmangle.SetComplement(
			channelAllColumns,
			channelPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ChannelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testChannelsUpsert(t *testing.T) {
	t.Parallel()

	if len(channelAllColumns) == len(channelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Channel{}
	if err = randomize.Struct(seed, &o, channelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Channel: %s", err)
	}

	count, err := Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, channelDBTypes, false, channelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Channel: %s", err)
	}

	count, err = Channels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}