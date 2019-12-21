package goarea

import "math"

// Pi proporçao numerica definida pela relação entre
// o perímetro de uma circunferencia e um diametro
const Pi = 3.1416

// Circ é responsável por calcular a área da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
 }
