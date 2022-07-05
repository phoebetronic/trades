package trades

func (t *Trade) Empty() bool {
	return t.PR == 0 && t.LO == 0 && t.SH == 0 && t.TS.AsTime().IsZero()
}
