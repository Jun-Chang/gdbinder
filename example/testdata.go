package example

import "database/sql"

//go:generate gdbinder -type=TestStruct
type TestStruct struct {
	ID    int `db:"id" foo:"var"`
	Col1  int `db:"col1"`
	Col2  int `db:"col2"`
	Col3  int `db:"col3"`
	Col4  int `db:"col4"`
	Col5  int `db:"col5"`
	Col6  int `db:"col6"`
	Col7  int `db:"col7"`
	Col8  int `db:"col8"`
	Col9  int `db:"col9"`
}

//go:generate gdbinder -type=TestStruct2
type TestStruct2 struct {
	ID     int `db:"id"`
	Col1   int `db:"col1"`
	Col2   int `db:"col2"`
	Col3   int `db:"col3"`
	Col4   int `db:"col4"`
	Col5   int `db:"col5"`
	Col6   int `db:"col6"`
	Col7   int `db:"col7"`
	Col8   int `db:"col8"`
	Col9   int `db:"col9"`
	Col10  int `db:"col10"`
	Col11  int `db:"col11"`
	Col12  int `db:"col12"`
	Col13  int `db:"col13"`
	Col14  int `db:"col14"`
	Col15  int `db:"col15"`
	Col16  int `db:"col16"`
	Col17  int `db:"col17"`
	Col18  int `db:"col18"`
	Col19  int `db:"col19"`
	Col20  int `db:"col20"`
	Col21  int `db:"col21"`
	Col22  int `db:"col22"`
	Col23  int `db:"col23"`
	Col24  int `db:"col24"`
	Col25  int `db:"col25"`
	Col26  int `db:"col26"`
	Col27  int `db:"col27"`
	Col28  int `db:"col28"`
	Col29  int `db:"col29"`
	Col30  int `db:"col30"`
	Col31  int `db:"col31"`
	Col32  int `db:"col32"`
	Col33  int `db:"col33"`
	Col34  int `db:"col34"`
	Col35  int `db:"col35"`
	Col36  int `db:"col36"`
	Col37  int `db:"col37"`
	Col38  int `db:"col38"`
	Col39  int `db:"col39"`
	Col40  int `db:"col40"`
	Col41  int `db:"col41"`
	Col42  int `db:"col42"`
	Col43  int `db:"col43"`
	Col44  int `db:"col44"`
	Col45  int `db:"col45"`
	Col46  int `db:"col46"`
	Col47  int `db:"col47"`
	Col48  int `db:"col48"`
	Col49  int `db:"col49"`
	Col50  int `db:"col50"`
	Col51  int `db:"col51"`
	Col52  int `db:"col52"`
	Col53  int `db:"col53"`
	Col54  int `db:"col54"`
	Col55  int `db:"col55"`
	Col56  int `db:"col56"`
	Col57  int `db:"col57"`
	Col58  int `db:"col58"`
	Col59  int `db:"col59"`
	Col60  int `db:"col60"`
	Col61  int `db:"col61"`
	Col62  int `db:"col62"`
	Col63  int `db:"col63"`
	Col64  int `db:"col64"`
	Col65  int `db:"col65"`
	Col66  int `db:"col66"`
	Col67  int `db:"col67"`
	Col68  int `db:"col68"`
	Col69  int `db:"col69"`
	Col70  int `db:"col70"`
	Col71  int `db:"col71"`
	Col72  int `db:"col72"`
	Col73  int `db:"col73"`
	Col74  int `db:"col74"`
	Col75  int `db:"col75"`
	Col76  int `db:"col76"`
	Col77  int `db:"col77"`
	Col78  int `db:"col78"`
	Col79  int `db:"col79"`
	Col80  int `db:"col80"`
	Col81  int `db:"col81"`
	Col82  int `db:"col82"`
	Col83  int `db:"col83"`
	Col84  int `db:"col84"`
	Col85  int `db:"col85"`
	Col86  int `db:"col86"`
	Col87  int `db:"col87"`
	Col88  int `db:"col88"`
	Col89  int `db:"col89"`
	Col90  int `db:"col90"`
	Col91  int `db:"col91"`
	Col92  int `db:"col92"`
	Col93  int `db:"col93"`
	Col94  int `db:"col94"`
	Col95  int `db:"col95"`
	Col96  int `db:"col96"`
	Col97  int `db:"col97"`
	Col98  int `db:"col98"`
	Col99  int `db:"col99"`
}

// FindStruct returns TestStruct slice (bound gdbinder_test table)
func FindStruct(db *sql.DB) ([]*TestStruct, error) {
	query := "SELECT * FROM gdbinder_test"
	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	return TestStructScan(rs)
}

// FindStruct2 returns TestStruct2 slice (bound gdbinder_test2 table)
func FindStruct2(db *sql.DB) ([]*TestStruct2, error) {
	query := "SELECT * FROM gdbinder_test2"
	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	return TestStruct2Scan(rs)
}
