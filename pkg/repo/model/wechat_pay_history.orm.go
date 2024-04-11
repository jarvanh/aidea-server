package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

}

// WechatPayHistoryN is a WechatPayHistory object, all fields are nullable
type WechatPayHistoryN struct {
	original              *wechatPayHistoryOriginal
	wechatPayHistoryModel *WechatPayHistoryModel

	Id          null.Int    `json:"id"`
	UserId      null.Int    `json:"user_id,omitempty"`
	PaymentId   null.String `json:"payment_id,omitempty"`
	ProductId   null.String `json:"product_id,omitempty"`
	Extra       null.String `json:"extra,omitempty"`
	Amount      null.Int    `json:"amount,omitempty"`
	Environment null.String `json:"environment"`
	Status      null.Int    `json:"status,omitempty"`
	PurchaseAt  null.Time   `json:"purchase_at"`
	Note        null.String `json:"note"`
	CreatedAt   null.Time
	UpdatedAt   null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *WechatPayHistoryN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for WechatPayHistory
func (inst *WechatPayHistoryN) SetModel(wechatPayHistoryModel *WechatPayHistoryModel) {
	inst.wechatPayHistoryModel = wechatPayHistoryModel
}

// wechatPayHistoryOriginal is an object which stores original WechatPayHistory from database
type wechatPayHistoryOriginal struct {
	Id          null.Int
	UserId      null.Int
	PaymentId   null.String
	ProductId   null.String
	Extra       null.String
	Amount      null.Int
	Environment null.String
	Status      null.Int
	PurchaseAt  null.Time
	Note        null.String
	CreatedAt   null.Time
	UpdatedAt   null.Time
}

// Staled identify whether the object has been modified
func (inst *WechatPayHistoryN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &wechatPayHistoryOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.UserId != inst.original.UserId {
			return true
		}
		if inst.PaymentId != inst.original.PaymentId {
			return true
		}
		if inst.ProductId != inst.original.ProductId {
			return true
		}
		if inst.Extra != inst.original.Extra {
			return true
		}
		if inst.Amount != inst.original.Amount {
			return true
		}
		if inst.Environment != inst.original.Environment {
			return true
		}
		if inst.Status != inst.original.Status {
			return true
		}
		if inst.PurchaseAt != inst.original.PurchaseAt {
			return true
		}
		if inst.Note != inst.original.Note {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "user_id":
				if inst.UserId != inst.original.UserId {
					return true
				}
			case "payment_id":
				if inst.PaymentId != inst.original.PaymentId {
					return true
				}
			case "product_id":
				if inst.ProductId != inst.original.ProductId {
					return true
				}
			case "extra":
				if inst.Extra != inst.original.Extra {
					return true
				}
			case "amount":
				if inst.Amount != inst.original.Amount {
					return true
				}
			case "environment":
				if inst.Environment != inst.original.Environment {
					return true
				}
			case "status":
				if inst.Status != inst.original.Status {
					return true
				}
			case "purchase_at":
				if inst.PurchaseAt != inst.original.PurchaseAt {
					return true
				}
			case "note":
				if inst.Note != inst.original.Note {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *WechatPayHistoryN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &wechatPayHistoryOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.UserId != inst.original.UserId {
			kv["user_id"] = inst.UserId
		}
		if inst.PaymentId != inst.original.PaymentId {
			kv["payment_id"] = inst.PaymentId
		}
		if inst.ProductId != inst.original.ProductId {
			kv["product_id"] = inst.ProductId
		}
		if inst.Extra != inst.original.Extra {
			kv["extra"] = inst.Extra
		}
		if inst.Amount != inst.original.Amount {
			kv["amount"] = inst.Amount
		}
		if inst.Environment != inst.original.Environment {
			kv["environment"] = inst.Environment
		}
		if inst.Status != inst.original.Status {
			kv["status"] = inst.Status
		}
		if inst.PurchaseAt != inst.original.PurchaseAt {
			kv["purchase_at"] = inst.PurchaseAt
		}
		if inst.Note != inst.original.Note {
			kv["note"] = inst.Note
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "user_id":
				if inst.UserId != inst.original.UserId {
					kv["user_id"] = inst.UserId
				}
			case "payment_id":
				if inst.PaymentId != inst.original.PaymentId {
					kv["payment_id"] = inst.PaymentId
				}
			case "product_id":
				if inst.ProductId != inst.original.ProductId {
					kv["product_id"] = inst.ProductId
				}
			case "extra":
				if inst.Extra != inst.original.Extra {
					kv["extra"] = inst.Extra
				}
			case "amount":
				if inst.Amount != inst.original.Amount {
					kv["amount"] = inst.Amount
				}
			case "environment":
				if inst.Environment != inst.original.Environment {
					kv["environment"] = inst.Environment
				}
			case "status":
				if inst.Status != inst.original.Status {
					kv["status"] = inst.Status
				}
			case "purchase_at":
				if inst.PurchaseAt != inst.original.PurchaseAt {
					kv["purchase_at"] = inst.PurchaseAt
				}
			case "note":
				if inst.Note != inst.original.Note {
					kv["note"] = inst.Note
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *WechatPayHistoryN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.wechatPayHistoryModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.wechatPayHistoryModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a wechat_pay_history
func (inst *WechatPayHistoryN) Delete(ctx context.Context) error {
	if inst.wechatPayHistoryModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.wechatPayHistoryModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *WechatPayHistoryN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type wechatPayHistoryScope struct {
	name  string
	apply func(builder query.Condition)
}

var wechatPayHistoryGlobalScopes = make([]wechatPayHistoryScope, 0)
var wechatPayHistoryLocalScopes = make([]wechatPayHistoryScope, 0)

// AddGlobalScopeForWechatPayHistory assign a global scope to a model
func AddGlobalScopeForWechatPayHistory(name string, apply func(builder query.Condition)) {
	wechatPayHistoryGlobalScopes = append(wechatPayHistoryGlobalScopes, wechatPayHistoryScope{name: name, apply: apply})
}

// AddLocalScopeForWechatPayHistory assign a local scope to a model
func AddLocalScopeForWechatPayHistory(name string, apply func(builder query.Condition)) {
	wechatPayHistoryLocalScopes = append(wechatPayHistoryLocalScopes, wechatPayHistoryScope{name: name, apply: apply})
}

func (m *WechatPayHistoryModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range wechatPayHistoryGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range wechatPayHistoryLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *WechatPayHistoryModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *WechatPayHistoryModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type WechatPayHistory struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id,omitempty"`
	PaymentId   string    `json:"payment_id,omitempty"`
	ProductId   string    `json:"product_id,omitempty"`
	Extra       string    `json:"extra,omitempty"`
	Amount      int64     `json:"amount,omitempty"`
	Environment string    `json:"environment"`
	Status      int       `json:"status,omitempty"`
	PurchaseAt  time.Time `json:"purchase_at"`
	Note        string    `json:"note"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (w WechatPayHistory) ToWechatPayHistoryN(allows ...string) WechatPayHistoryN {
	if len(allows) == 0 {
		return WechatPayHistoryN{

			Id:          null.IntFrom(int64(w.Id)),
			UserId:      null.IntFrom(int64(w.UserId)),
			PaymentId:   null.StringFrom(w.PaymentId),
			ProductId:   null.StringFrom(w.ProductId),
			Extra:       null.StringFrom(w.Extra),
			Amount:      null.IntFrom(int64(w.Amount)),
			Environment: null.StringFrom(w.Environment),
			Status:      null.IntFrom(int64(w.Status)),
			PurchaseAt:  null.TimeFrom(w.PurchaseAt),
			Note:        null.StringFrom(w.Note),
			CreatedAt:   null.TimeFrom(w.CreatedAt),
			UpdatedAt:   null.TimeFrom(w.UpdatedAt),
		}
	}

	res := WechatPayHistoryN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "user_id":
			res.UserId = null.IntFrom(int64(w.UserId))
		case "payment_id":
			res.PaymentId = null.StringFrom(w.PaymentId)
		case "product_id":
			res.ProductId = null.StringFrom(w.ProductId)
		case "extra":
			res.Extra = null.StringFrom(w.Extra)
		case "amount":
			res.Amount = null.IntFrom(int64(w.Amount))
		case "environment":
			res.Environment = null.StringFrom(w.Environment)
		case "status":
			res.Status = null.IntFrom(int64(w.Status))
		case "purchase_at":
			res.PurchaseAt = null.TimeFrom(w.PurchaseAt)
		case "note":
			res.Note = null.StringFrom(w.Note)
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w WechatPayHistory) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *WechatPayHistoryN) ToWechatPayHistory() WechatPayHistory {
	return WechatPayHistory{

		Id:          w.Id.Int64,
		UserId:      w.UserId.Int64,
		PaymentId:   w.PaymentId.String,
		ProductId:   w.ProductId.String,
		Extra:       w.Extra.String,
		Amount:      w.Amount.Int64,
		Environment: w.Environment.String,
		Status:      int(w.Status.Int64),
		PurchaseAt:  w.PurchaseAt.Time,
		Note:        w.Note.String,
		CreatedAt:   w.CreatedAt.Time,
		UpdatedAt:   w.UpdatedAt.Time,
	}
}

// WechatPayHistoryModel is a model which encapsulates the operations of the object
type WechatPayHistoryModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var wechatPayHistoryTableName = "wechat_pay_history"

// WechatPayHistoryTable return table name for WechatPayHistory
func WechatPayHistoryTable() string {
	return wechatPayHistoryTableName
}

const (
	FieldWechatPayHistoryId          = "id"
	FieldWechatPayHistoryUserId      = "user_id"
	FieldWechatPayHistoryPaymentId   = "payment_id"
	FieldWechatPayHistoryProductId   = "product_id"
	FieldWechatPayHistoryExtra       = "extra"
	FieldWechatPayHistoryAmount      = "amount"
	FieldWechatPayHistoryEnvironment = "environment"
	FieldWechatPayHistoryStatus      = "status"
	FieldWechatPayHistoryPurchaseAt  = "purchase_at"
	FieldWechatPayHistoryNote        = "note"
	FieldWechatPayHistoryCreatedAt   = "created_at"
	FieldWechatPayHistoryUpdatedAt   = "updated_at"
)

// WechatPayHistoryFields return all fields in WechatPayHistory model
func WechatPayHistoryFields() []string {
	return []string{
		"id",
		"user_id",
		"payment_id",
		"product_id",
		"extra",
		"amount",
		"environment",
		"status",
		"purchase_at",
		"note",
		"created_at",
		"updated_at",
	}
}

func SetWechatPayHistoryTable(tableName string) {
	wechatPayHistoryTableName = tableName
}

// NewWechatPayHistoryModel create a WechatPayHistoryModel
func NewWechatPayHistoryModel(db query.Database) *WechatPayHistoryModel {
	return &WechatPayHistoryModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           wechatPayHistoryTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *WechatPayHistoryModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *WechatPayHistoryModel) clone() *WechatPayHistoryModel {
	return &WechatPayHistoryModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *WechatPayHistoryModel) WithoutGlobalScopes(names ...string) *WechatPayHistoryModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *WechatPayHistoryModel) WithLocalScopes(names ...string) *WechatPayHistoryModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *WechatPayHistoryModel) Condition(builder query.SQLBuilder) *WechatPayHistoryModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *WechatPayHistoryModel) Find(ctx context.Context, id int64) (*WechatPayHistoryN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *WechatPayHistoryModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *WechatPayHistoryModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *WechatPayHistoryModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]WechatPayHistoryN, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(ctx, builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(ctx, append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *WechatPayHistoryModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]WechatPayHistoryN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"user_id",
			"payment_id",
			"product_id",
			"extra",
			"amount",
			"environment",
			"status",
			"purchase_at",
			"note",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "user_id":
			selectFields = append(selectFields, f)
		case "payment_id":
			selectFields = append(selectFields, f)
		case "product_id":
			selectFields = append(selectFields, f)
		case "extra":
			selectFields = append(selectFields, f)
		case "amount":
			selectFields = append(selectFields, f)
		case "environment":
			selectFields = append(selectFields, f)
		case "status":
			selectFields = append(selectFields, f)
		case "purchase_at":
			selectFields = append(selectFields, f)
		case "note":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*WechatPayHistoryN, []interface{}) {
		var wechatPayHistoryVar WechatPayHistoryN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &wechatPayHistoryVar.Id)
			case "user_id":
				scanFields = append(scanFields, &wechatPayHistoryVar.UserId)
			case "payment_id":
				scanFields = append(scanFields, &wechatPayHistoryVar.PaymentId)
			case "product_id":
				scanFields = append(scanFields, &wechatPayHistoryVar.ProductId)
			case "extra":
				scanFields = append(scanFields, &wechatPayHistoryVar.Extra)
			case "amount":
				scanFields = append(scanFields, &wechatPayHistoryVar.Amount)
			case "environment":
				scanFields = append(scanFields, &wechatPayHistoryVar.Environment)
			case "status":
				scanFields = append(scanFields, &wechatPayHistoryVar.Status)
			case "purchase_at":
				scanFields = append(scanFields, &wechatPayHistoryVar.PurchaseAt)
			case "note":
				scanFields = append(scanFields, &wechatPayHistoryVar.Note)
			case "created_at":
				scanFields = append(scanFields, &wechatPayHistoryVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &wechatPayHistoryVar.UpdatedAt)
			}
		}

		return &wechatPayHistoryVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	wechatPayHistorys := make([]WechatPayHistoryN, 0)
	for rows.Next() {
		wechatPayHistoryReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		wechatPayHistoryReal.original = &wechatPayHistoryOriginal{}
		_ = query.Copy(wechatPayHistoryReal, wechatPayHistoryReal.original)

		wechatPayHistoryReal.SetModel(m)
		wechatPayHistorys = append(wechatPayHistorys, *wechatPayHistoryReal)
	}

	return wechatPayHistorys, nil
}

// First return first result for given query
func (m *WechatPayHistoryModel) First(ctx context.Context, builders ...query.SQLBuilder) (*WechatPayHistoryN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new wechat_pay_history to database
func (m *WechatPayHistoryModel) Create(ctx context.Context, kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all wechat_pay_historys to database
func (m *WechatPayHistoryModel) SaveAll(ctx context.Context, wechatPayHistorys []WechatPayHistoryN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, wechatPayHistory := range wechatPayHistorys {
		id, err := m.Save(ctx, wechatPayHistory)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a wechat_pay_history to database
func (m *WechatPayHistoryModel) Save(ctx context.Context, wechatPayHistory WechatPayHistoryN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, wechatPayHistory.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new wechat_pay_history or update it when it has a id > 0
func (m *WechatPayHistoryModel) SaveOrUpdate(ctx context.Context, wechatPayHistory WechatPayHistoryN, onlyFields ...string) (id int64, updated bool, err error) {
	if wechatPayHistory.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, wechatPayHistory.Id.Int64, wechatPayHistory, onlyFields...)
		return wechatPayHistory.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, wechatPayHistory, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *WechatPayHistoryModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *WechatPayHistoryModel) Update(ctx context.Context, builder query.SQLBuilder, wechatPayHistory WechatPayHistoryN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, wechatPayHistory.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *WechatPayHistoryModel) UpdateById(ctx context.Context, id int64, wechatPayHistory WechatPayHistoryN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, wechatPayHistory.StaledKV(onlyFields...))
}

// Delete remove a model
func (m *WechatPayHistoryModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *WechatPayHistoryModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}
