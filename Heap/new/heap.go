package main


func main() {

}

type maax struct {
	maax []int
}

func (m *maax) heapfy(array []int) {

	m.maax = append(m.maax, array...)

	paridx := m.paridx(len(m.maax) - 1)

	for paridx >= 0 {
		m.swiftd(paridx)
		paridx--
	}
}
func (m *maax) isert(data int) {
	
      m.maax=append(m.maax, data)

	  m.swiftup(len(m.maax)-1)

}

func (m *maax) paridx(chilidx int) int {
	return (chilidx - 1) / 2
}

func (m *maax) lchildx(paridx int) int {
	return (paridx * 2) + 1
}

func (m *maax) rchildx(paridx int) int {
	return (paridx * 2) + 2
}

func (m *maax) swiftd(paridx int) {

	lchildx:=m.lchildx(paridx)
	endidx:=len(m.maax)-1

	for lchildx<=endidx{
		rchildx:=m.rchildx(paridx)

		idswft:=lchildx

		if rchildx<=endidx&&m.maax[rchildx]>m.maax[rchildx]{
			idswft=rchildx
		}

		if m.maax[paridx]>m.maax[idswft]{
			return
		}else {
			m.maax[paridx],m.maax[idswft]=m.maax[idswft],m.maax[paridx]
		}
        m.maax[paridx]=m.maax[idswft]
		

		lchildx=m.lchildx(paridx)

		
	}

}

func (m *maax) swiftup(chilidx int) {
	paridx := m.paridx(chilidx)

	for paridx >= 0 && m.maax[paridx] < m.maax[chilidx] {
		m.maax[paridx], m.maax[chilidx] = m.maax[chilidx], m.maax[paridx]

		chilidx = paridx
		paridx = m.paridx(chilidx)
	}
}
