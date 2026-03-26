package tests

import (
	"testing"

	"github.com/atticus64/dona/cmd/actions"
)

func TestSearchSort(t *testing.T) {
	query := "fedora"

	dots, err := actions.SearchDotfiles(query, 1)

	if err != nil {
		t.Errorf("Error al buscar dotfiles %s", err)
	}

	if len(dots) < 1 {
		t.Errorf("Deberia listar por lo menos %d repositorio, obtenidos %d", 1, len(dots))
	}

	actions.SortViaStars(dots)

	first := dots[0]
	second := dots[1]

	if first.Stargazers_count < second.Stargazers_count {
		t.Errorf("Las estrellas del primero son %d y del segundo %d, deberia estar ordenadas de mayor a menor", first.Stargazers_count, second.Stargazers_count)
	}

}
