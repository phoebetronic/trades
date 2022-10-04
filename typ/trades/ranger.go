package trades

type Ranger struct {
	off []*Trades
	win []*Trades
	rig []*Trades
}

func (r *Ranger) Off() []*Trades {
	return r.off
}

func (r *Ranger) Fir() *Trades {
	return r.win[len(r.win)-1]
}

func (r *Ranger) Win() []*Trades {
	return r.win
}

func (r *Ranger) Rig() []*Trades {
	return r.rig
}

func (r *Ranger) All() []*Trades {
	var all []*Trades

	{
		all = append(all, r.off...)
		all = append(all, r.win...)
		all = append(all, r.rig...)
	}

	return all
}
