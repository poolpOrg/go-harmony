package octaves

import (
	"testing"
)

func TestOctaves_length(t *testing.T) {
	octavesList := Octaves()
	if len(octavesList) != 8 {
		t.Fatalf(`octaves.Octaves() = %d, want %d`, len(octavesList), 8)
	}
}

func TestOctaves_Name(t *testing.T) {
	if C0.Name() != "C0" {
		t.Fatalf(`octaves.C0.Name() = %s, want %s`, C0.Name(), "C0")
	}
	if C1.Name() != "C1" {
		t.Fatalf(`octaves.C1.Name() = %s, want %s`, C1.Name(), "C1")
	}
	if C2.Name() != "C2" {
		t.Fatalf(`octaves.C2.Name() = %s, want %s`, C2.Name(), "C2")
	}
	if C3.Name() != "C3" {
		t.Fatalf(`octaves.C3.Name() = %s, want %s`, C3.Name(), "C3")
	}
	if C4.Name() != "C4" {
		t.Fatalf(`octaves.C4.Name() = %s, want %s`, C4.Name(), "C4")
	}
	if C5.Name() != "C5" {
		t.Fatalf(`octaves.C5.Name() = %s, want %s`, C5.Name(), "C5")
	}
	if C6.Name() != "C6" {
		t.Fatalf(`octaves.C6.Name() = %s, want %s`, C6.Name(), "C6")
	}
	if C7.Name() != "C7" {
		t.Fatalf(`octaves.C7.Name() = %s, want %s`, C7.Name(), "C7")
	}
}

func TestOctaves_Position(t *testing.T) {
	if C0.Position() != 0 {
		t.Fatalf(`octaves.C0.Position() = %d, want %d`, C0.Position(), 0)
	}
	if C1.Position() != 1 {
		t.Fatalf(`octaves.C1.Position() = %d, want %d`, C1.Position(), 1)
	}
	if C2.Position() != 2 {
		t.Fatalf(`octaves.C2.Position() = %d, want %d`, C2.Position(), 2)
	}
	if C3.Position() != 3 {
		t.Fatalf(`octaves.C3.Position() = %d, want %d`, C3.Position(), 3)
	}
	if C4.Position() != 4 {
		t.Fatalf(`octaves.C4.Position() = %d, want %d`, C4.Position(), 4)
	}
	if C5.Position() != 5 {
		t.Fatalf(`octaves.C5.Position() = %d, want %d`, C5.Position(), 5)
	}
	if C6.Position() != 6 {
		t.Fatalf(`octaves.C6.Position() = %d, want %d`, C6.Position(), 6)
	}
	if C7.Position() != 7 {
		t.Fatalf(`octaves.C7.Position() = %d, want %d`, C7.Position(), 7)
	}
}

func TestOctaves_Previous(t *testing.T) {
	if C0.Previous() != nil {
		t.Fatalf(`octaves.C0.Previous() = %v, want %v`, C0.Previous().Name(), nil)
	}
	if *C1.Previous() != C0 {
		t.Fatalf(`octaves.C1.Previous() = %v, want %v`, C1.Previous().Name(), C0.Name())
	}
	if *C2.Previous() != C1 {
		t.Fatalf(`octaves.C2.Previous() = %v, want %v`, C2.Previous().Name(), C1.Name())
	}
	if *C3.Previous() != C2 {
		t.Fatalf(`octaves.C3.Previous() = %v, want %v`, C3.Previous().Name(), C2.Name())
	}
	if *C4.Previous() != C3 {
		t.Fatalf(`octaves.C4.Previous() = %v, want %v`, C4.Previous().Name(), C3.Name())
	}
	if *C5.Previous() != C4 {
		t.Fatalf(`octaves.C5.Previous() = %v, want %v`, C5.Previous().Name(), C4.Name())
	}
	if *C6.Previous() != C5 {
		t.Fatalf(`octaves.C6.Previous() = %v, want %v`, C6.Previous().Name(), C5.Name())
	}
	if *C7.Previous() != C6 {
		t.Fatalf(`octaves.C7.Previous() = %v, want %v`, C7.Previous().Name(), C6.Name())
	}
}

func TestOctaves_Next(t *testing.T) {
	if *C0.Next() != C1 {
		t.Fatalf(`octaves.C0.Next() = %v, want %v`, C0.Next().Name(), C1.Name())
	}
	if *C1.Next() != C2 {
		t.Fatalf(`octaves.C1.Next() = %v, want %v`, C1.Next().Name(), C2.Name())
	}
	if *C2.Next() != C3 {
		t.Fatalf(`octaves.C2.Next() = %v, want %v`, C2.Next().Name(), C3.Name())
	}
	if *C3.Next() != C4 {
		t.Fatalf(`octaves.C3.Next() = %v, want %v`, C3.Next().Name(), C4.Name())
	}
	if *C4.Next() != C5 {
		t.Fatalf(`octaves.C4.Next() = %v, want %v`, C4.Next().Name(), C5.Name())
	}
	if *C5.Next() != C6 {
		t.Fatalf(`octaves.C5.Next() = %v, want %v`, C5.Next().Name(), C6.Name())
	}
	if *C6.Next() != C7 {
		t.Fatalf(`octaves.C6.Next() = %v, want %v`, C6.Next().Name(), C7.Name())
	}
	if C7.Next() != nil {
		t.Fatalf(`octaves.C7.Next() = %v, want %v`, C7.Next().Name(), nil)
	}
}

func TestOctaves_Add(t *testing.T) {
	if *C0.Add(1) != C1 {
		t.Fatalf(`octaves.C0.Add(1) = %v, want %v`, C0.Next().Name(), C1.Name())
	}
	if *C0.Add(2) != C2 {
		t.Fatalf(`octaves.C0.Add(2) = %v, want %v`, C0.Next().Name(), C2.Name())
	}
	if *C0.Add(3) != C3 {
		t.Fatalf(`octaves.C0.Add(3) = %v, want %v`, C0.Next().Name(), C3.Name())
	}
	if *C0.Add(4) != C4 {
		t.Fatalf(`octaves.C0.Add(4) = %v, want %v`, C0.Next().Name(), C4.Name())
	}
	if *C0.Add(5) != C5 {
		t.Fatalf(`octaves.C0.Add(5) = %v, want %v`, C0.Next().Name(), C5.Name())
	}
	if *C0.Add(6) != C6 {
		t.Fatalf(`octaves.C0.Add(6) = %v, want %v`, C0.Next().Name(), C6.Name())
	}
	if *C0.Add(7) != C7 {
		t.Fatalf(`octaves.C0.Add(7) = %v, want %v`, C0.Next().Name(), C7.Name())
	}
	if C0.Add(8) != nil {
		t.Fatalf(`octaves.C0.Add(8) = %v, want %v`, C0.Next().Name(), nil)
	}

}

func TestOctaves_Substract(t *testing.T) {
	if *C7.Substract(1) != C6 {
		t.Fatalf(`octaves.C7.Substract(1) = %v, want %v`, C7.Next().Name(), C6.Name())
	}
	if *C7.Substract(2) != C5 {
		t.Fatalf(`octaves.C7.Substract(2) = %v, want %v`, C7.Next().Name(), C5.Name())
	}
	if *C7.Substract(3) != C4 {
		t.Fatalf(`octaves.C7.Substract(3) = %v, want %v`, C7.Next().Name(), C4.Name())
	}
	if *C7.Substract(4) != C3 {
		t.Fatalf(`octaves.C7.Substract(4) = %v, want %v`, C7.Next().Name(), C3.Name())
	}
	if *C7.Substract(5) != C2 {
		t.Fatalf(`octaves.C7.Substract(5) = %v, want %v`, C7.Next().Name(), C2.Name())
	}
	if *C7.Substract(6) != C1 {
		t.Fatalf(`octaves.C7.Substract(6) = %v, want %v`, C7.Next().Name(), C1.Name())
	}
	if *C7.Substract(7) != C0 {
		t.Fatalf(`octaves.C7.Substract(7) = %v, want %v`, C7.Next().Name(), C0.Name())
	}
	if C7.Substract(8) != nil {
		t.Fatalf(`octaves.C7.Substract(8) = %v, want %v`, C7.Next().Name(), nil)
	}
}

func TestNaturals_FromName(t *testing.T) {
	if octave, err := FromName("C0"); err != nil {
		t.Fatalf(`octaves.FromName("C0") = error: %s`, err)
	} else {
		if *octave != C0 {
			t.Fatalf(`octaves.FromName("C0") = %v, want %v`, octave.Name(), C0.Name())
		}
	}

	if octave, err := FromName("C1"); err != nil {
		t.Fatalf(`octaves.FromName("C1") = error: %s`, err)
	} else {
		if *octave != C1 {
			t.Fatalf(`octaves.FromName("C1") = %v, want %v`, octave.Name(), C1.Name())
		}
	}

	if octave, err := FromName("C2"); err != nil {
		t.Fatalf(`octaves.FromName("C2") = error: %s`, err)
	} else {
		if *octave != C2 {
			t.Fatalf(`octaves.FromName("C2") = %v, want %v`, octave.Name(), C2.Name())
		}
	}

	if octave, err := FromName("C3"); err != nil {
		t.Fatalf(`octaves.FromName("C3") = error: %s`, err)
	} else {
		if *octave != C3 {
			t.Fatalf(`octaves.FromName("C3") = %v, want %v`, octave.Name(), C3.Name())
		}
	}

	if octave, err := FromName("C4"); err != nil {
		t.Fatalf(`octaves.FromName("C4") = error: %s`, err)
	} else {
		if *octave != C4 {
			t.Fatalf(`octaves.FromName("C4") = %v, want %v`, octave.Name(), C4.Name())
		}
	}

	if octave, err := FromName("C5"); err != nil {
		t.Fatalf(`octaves.FromName("C5") = error: %s`, err)
	} else {
		if *octave != C5 {
			t.Fatalf(`octaves.FromName("C5") = %v, want %v`, octave.Name(), C5.Name())
		}
	}

	if octave, err := FromName("C6"); err != nil {
		t.Fatalf(`octaves.FromName("C6") = error: %s`, err)
	} else {
		if *octave != C6 {
			t.Fatalf(`octaves.FromName("C6") = %v, want %v`, octave.Name(), C6.Name())
		}
	}

	if octave, err := FromName("C7"); err != nil {
		t.Fatalf(`octaves.FromName("C7") = error: %s`, err)
	} else {
		if *octave != C7 {
			t.Fatalf(`octaves.FromName("C7") = %v, want %v`, octave.Name(), C7.Name())
		}
	}
}

func TestNaturals_FromPosition(t *testing.T) {
	if octave, err := FromPosition(0); err != nil {
		t.Fatalf(`octaves.FromPosition(0) = error: %s`, err)
	} else {
		if *octave != C0 {
			t.Fatalf(`octaves.FromPosition(0) = %v, want %v`, octave.Name(), C0.Name())
		}
	}

	if octave, err := FromPosition(1); err != nil {
		t.Fatalf(`octaves.FromPosition(1) = error: %s`, err)
	} else {
		if *octave != C1 {
			t.Fatalf(`octaves.FromPosition(1) = %v, want %v`, octave.Name(), C1.Name())
		}
	}

	if octave, err := FromPosition(2); err != nil {
		t.Fatalf(`octaves.FromPosition(2) = error: %s`, err)
	} else {
		if *octave != C2 {
			t.Fatalf(`octaves.FromPosition(2) = %v, want %v`, octave.Name(), C2.Name())
		}
	}

	if octave, err := FromPosition(3); err != nil {
		t.Fatalf(`octaves.FromPosition(3) = error: %s`, err)
	} else {
		if *octave != C3 {
			t.Fatalf(`octaves.FromPosition(3) = %v, want %v`, octave.Name(), C3.Name())
		}
	}

	if octave, err := FromPosition(4); err != nil {
		t.Fatalf(`octaves.FromPosition(4) = error: %s`, err)
	} else {
		if *octave != C4 {
			t.Fatalf(`octaves.FromPosition(4) = %v, want %v`, octave.Name(), C4.Name())
		}
	}

	if octave, err := FromPosition(5); err != nil {
		t.Fatalf(`octaves.FromPosition(5) = error: %s`, err)
	} else {
		if *octave != C5 {
			t.Fatalf(`octaves.FromPosition(5) = %v, want %v`, octave.Name(), C5.Name())
		}
	}

	if octave, err := FromPosition(6); err != nil {
		t.Fatalf(`octaves.FromPosition(6) = error: %s`, err)
	} else {
		if *octave != C6 {
			t.Fatalf(`octaves.FromPosition(6) = %v, want %v`, octave.Name(), C6.Name())
		}
	}

	if octave, err := FromPosition(7); err != nil {
		t.Fatalf(`octaves.FromPosition(7) = error: %s`, err)
	} else {
		if *octave != C7 {
			t.Fatalf(`octaves.FromPosition(7) = %v, want %v`, octave.Name(), C7.Name())
		}
	}
}
