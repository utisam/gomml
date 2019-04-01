package examples

import "time"

//go:generate mockgen -package ${GOPACKAGE} -source ${GOFILE} -destination mock_${GOFILE}

// ExampleInterface is a interface of example to be tested
type ExampleInterface interface {
	Create(
		name string,
		hashedPassword []byte,
		createdAt, updatedAt time.Time,
	)
}
