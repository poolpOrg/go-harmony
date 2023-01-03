package naturals

import (
	"testing"
)

func TestNaturals_length(t *testing.T) {
	naturalsList := Naturals()
	if len(naturalsList) != 7 {
		t.Fatalf(`naturals.Naturals() = %d, want %d`, len(naturalsList), 7)
	}
}

func TestNaturals_Name(t *testing.T) {
	if C.Name() != "C" {
		t.Fatalf(`naturals.C.Name() = %s, want %s`, C.Name(), "C")
	}
	if D.Name() != "D" {
		t.Fatalf(`naturals.D.Name() = %s, want %s`, D.Name(), "D")
	}
	if E.Name() != "E" {
		t.Fatalf(`naturals.E.Name() = %s, want %s`, E.Name(), "E")
	}
	if F.Name() != "F" {
		t.Fatalf(`naturals.F.Name() = %s, want %s`, F.Name(), "F")
	}
	if G.Name() != "G" {
		t.Fatalf(`naturals.G.Name() = %s, want %s`, G.Name(), "G")
	}
	if A.Name() != "A" {
		t.Fatalf(`naturals.A.Name() = %s, want %s`, A.Name(), "A")
	}
	if B.Name() != "B" {
		t.Fatalf(`naturals.B.Name() = %s, want %s`, B.Name(), "B")
	}
}

func TestNaturals_Position(t *testing.T) {
	if C.Position() != 0 {
		t.Fatalf(`naturals.C.Position() = %d, want %d`, C.Position(), 0)
	}
	if D.Position() != 1 {
		t.Fatalf(`naturals.D.Position() = %d, want %d`, D.Position(), 1)
	}
	if E.Position() != 2 {
		t.Fatalf(`naturals.E.Position() = %d, want %d`, E.Position(), 2)
	}
	if F.Position() != 3 {
		t.Fatalf(`naturals.F.Position() = %d, want %d`, F.Position(), 3)
	}
	if G.Position() != 4 {
		t.Fatalf(`naturals.G.Position() = %d, want %d`, G.Position(), 4)
	}
	if A.Position() != 5 {
		t.Fatalf(`naturals.A.Position() = %d, want %d`, A.Position(), 5)
	}
	if B.Position() != 6 {
		t.Fatalf(`naturals.B.Position() = %d, want %d`, B.Position(), 6)
	}
}

func TestNaturals_Semitones(t *testing.T) {
	if C.Semitones() != 0 {
		t.Fatalf(`naturals.C.Semitones() = %d, want %d`, C.Semitones(), 0)
	}
	if D.Semitones() != 2 {
		t.Fatalf(`naturals.D.Semitones() = %d, want %d`, D.Semitones(), 2)
	}
	if E.Semitones() != 4 {
		t.Fatalf(`naturals.E.Semitones() = %d, want %d`, E.Semitones(), 4)
	}
	if F.Semitones() != 5 {
		t.Fatalf(`naturals.F.Semitones() = %d, want %d`, F.Semitones(), 5)
	}
	if G.Semitones() != 7 {
		t.Fatalf(`naturals.G.Semitones() = %d, want %d`, G.Semitones(), 7)
	}
	if A.Semitones() != 9 {
		t.Fatalf(`naturals.A.Semitones() = %d, want %d`, A.Semitones(), 9)
	}
	if B.Semitones() != 11 {
		t.Fatalf(`naturals.B.Semitones() = %d, want %d`, B.Semitones(), 11)
	}
}

func TestNaturals_Previous(t *testing.T) {
	if *C.Previous() != B {
		t.Fatalf(`naturals.C.Previous() = %v, want %v`, C.Previous(), B)
	}
	if *D.Previous() != C {
		t.Fatalf(`naturals.D.Previous() = %v, want %v`, D.Previous(), C)
	}
	if *E.Previous() != D {
		t.Fatalf(`naturals.E.Previous() = %v, want %v`, E.Previous(), D)
	}
	if *F.Previous() != E {
		t.Fatalf(`naturals.F.Previous() = %v, want %v`, F.Previous(), E)
	}
	if *G.Previous() != F {
		t.Fatalf(`naturals.G.Previous() = %v, want %v`, G.Previous(), F)
	}
	if *A.Previous() != G {
		t.Fatalf(`naturals.A.Previous() = %v, want %v`, A.Previous(), G)
	}
	if *B.Previous() != A {
		t.Fatalf(`naturals.B.Previous() = %v, want %v`, B.Previous(), A)
	}
}

func TestNaturals_Next(t *testing.T) {
	if *C.Next() != D {
		t.Fatalf(`naturals.C.Next() = %v, want %v`, C.Next(), D)
	}
	if *D.Next() != E {
		t.Fatalf(`naturals.D.Next() = %v, want %v`, D.Next(), E)
	}
	if *E.Next() != F {
		t.Fatalf(`naturals.E.Next() = %v, want %v`, E.Next(), F)
	}
	if *F.Next() != G {
		t.Fatalf(`naturals.F.Next() = %v, want %v`, F.Next(), G)
	}
	if *G.Next() != A {
		t.Fatalf(`naturals.G.Next() = %v, want %v`, G.Next(), A)
	}
	if *A.Next() != B {
		t.Fatalf(`naturals.A.Next() = %v, want %v`, A.Next(), B)
	}
	if *B.Next() != C {
		t.Fatalf(`naturals.B.Next() = %v, want %v`, B.Next(), C)
	}
}

func TestNaturals_FromName(t *testing.T) {
	if natural, err := FromName("C"); err != nil {
		t.Fatalf(`naturals.FromName("C") = error: %s`, err)
	} else {
		if *natural != C {
			t.Fatalf(`naturals.FromName("C") = %v, want %v`, natural.Name(), C.Name())
		}
	}

	if natural, err := FromName("D"); err != nil {
		t.Fatalf(`naturals.FromName("D") = error: %s`, err)
	} else {
		if *natural != D {
			t.Fatalf(`naturals.FromName("D") = %v, want %v`, natural.Name(), D.Name())
		}
	}

	if natural, err := FromName("E"); err != nil {
		t.Fatalf(`naturals.FromName("E") = error: %s`, err)
	} else {
		if *natural != E {
			t.Fatalf(`naturals.FromName("E") = %v, want %v`, natural.Name(), E.Name())
		}
	}

	if natural, err := FromName("F"); err != nil {
		t.Fatalf(`naturals.FromName("F") = error: %s`, err)
	} else {
		if *natural != F {
			t.Fatalf(`naturals.FromName("F") = %v, want %v`, natural.Name(), F.Name())
		}
	}

	if natural, err := FromName("G"); err != nil {
		t.Fatalf(`naturals.FromName("G") = error: %s`, err)
	} else {
		if *natural != G {
			t.Fatalf(`naturals.FromName("G") = %v, want %v`, natural.Name(), G.Name())
		}
	}

	if natural, err := FromName("A"); err != nil {
		t.Fatalf(`naturals.FromName("A") = error: %s`, err)
	} else {
		if *natural != A {
			t.Fatalf(`naturals.FromName("A") = %v, want %v`, natural.Name(), A.Name())
		}
	}

	if natural, err := FromName("B"); err != nil {
		t.Fatalf(`naturals.FromName("B") = error: %s`, err)
	} else {
		if *natural != B {
			t.Fatalf(`naturals.FromName("B") = %v, want %v`, natural.Name(), B.Name())
		}
	}

	if natural, err := FromName("Z"); err == nil {
		t.Fatalf(`naturals.FromName("B") = %v, want %v`, natural, nil)
	}
}
