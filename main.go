package main

import (
	"fmt"
	"sync"
	"time"
)

/*
PROBLEMA: Sistema de Rate Limiting con Ventana Deslizante

Implementa un sistema de rate limiting que controle el número de requests
por usuario usando una ventana de tiempo deslizante (sliding window).

REQUISITOS:
1. Limitar a N requests por usuario en los últimos M segundos
2. Método IsAllowed(userID string) bool - retorna si el request es permitido
3. Método Reset(userID string) - limpia el historial de un usuario
4. El sistema debe ser thread-safe (manejar concurrencia)
5. Optimizar el uso de memoria (limpiar datos antiguos)

EJEMPLO:
- Límite: 5 requests en 10 segundos
- Usuario "user1" hace 5 requests -> todos permitidos
- Usuario "user1" hace 1 request más inmediatamente -> bloqueado
- Después de 10 segundos, puede hacer requests nuevamente

BONUS (si da tiempo):
- Implementar GetStats() que retorne estadísticas por usuario
- Agregar diferentes límites por tipo de usuario (free/premium)
- Implementar un método para limpiar usuarios inactivos
*/

type RateLimiter struct {
	maxRequests    int
	windowDuration time.Duration
	// TODO: Agregar las estructuras de datos necesarias
	// Piensa en: ¿Cómo almacenar timestamps? ¿Cómo manejar múltiples usuarios?
	// ¿Cómo hacer thread-safe? ¿Cómo limpiar datos viejos eficientemente?

	mu sync.Mutex
}

// NewRateLimiter crea una nueva instancia del rate limiter
func NewRateLimiter(maxRequests int, windowDuration time.Duration) *RateLimiter {
	// TODO: Implementar constructor
	return &RateLimiter{
		maxRequests:    maxRequests,
		windowDuration: windowDuration,
	}
}

// IsAllowed verifica si un usuario puede hacer un request
func (rl *RateLimiter) IsAllowed(userID string) bool {
	// TODO: Implementar lógica de rate limiting
	// 1. Obtener historial del usuario
	// 2. Limpiar requests fuera de la ventana de tiempo
	// 3. Verificar si puede hacer el request
	// 4. Si es permitido, agregar el timestamp actual
	// 5. Retornar resultado

	return false
}

// Reset limpia el historial de requests de un usuario
func (rl *RateLimiter) Reset(userID string) {
	// TODO: Implementar reset para un usuario específico
}

// GetStats retorna estadísticas del rate limiter (BONUS)
func (rl *RateLimiter) GetStats(userID string) map[string]interface{} {
	// TODO: Retornar información útil como:
	// - Requests en la ventana actual
	// - Tiempo hasta el próximo request disponible
	// - Total de requests bloqueados
	return nil
}

// CleanupInactiveUsers elimina usuarios sin actividad reciente (BONUS)
func (rl *RateLimiter) CleanupInactiveUsers(inactivityThreshold time.Duration) int {
	// TODO: Limpiar usuarios que no han hecho requests recientemente
	// Retornar número de usuarios eliminados
	return 0
}

// ==================== CASOS DE PRUEBA ====================

func main() {
	fmt.Println("=== Test 1: Funcionamiento Básico ===")
	testBasicFunctionality()

	fmt.Println("\n=== Test 2: Ventana Deslizante ===")
	testSlidingWindow()

	fmt.Println("\n=== Test 3: Múltiples Usuarios ===")
	testMultipleUsers()

	fmt.Println("\n=== Test 4: Concurrencia ===")
	testConcurrency()
}

func testBasicFunctionality() {
	rl := NewRateLimiter(3, 5*time.Second)

	// Los primeros 3 requests deben ser permitidos
	for i := 1; i <= 3; i++ {
		if !rl.IsAllowed("user1") {
			fmt.Printf("❌ Request %d debería ser permitido\n", i)
		} else {
			fmt.Printf("✓ Request %d permitido\n", i)
		}
	}

	// El 4to request debe ser bloqueado
	if rl.IsAllowed("user1") {
		fmt.Println("❌ Request 4 debería ser bloqueado")
	} else {
		fmt.Println("✓ Request 4 bloqueado correctamente")
	}
}

func testSlidingWindow() {
	rl := NewRateLimiter(2, 2*time.Second)

	// Hacer 2 requests
	rl.IsAllowed("user2")
	rl.IsAllowed("user2")

	// Esperar 1 segundo
	time.Sleep(1 * time.Second)

	// Este debe ser bloqueado (aún dentro de la ventana)
	if rl.IsAllowed("user2") {
		fmt.Println("❌ Request debería ser bloqueado dentro de la ventana")
	} else {
		fmt.Println("✓ Request bloqueado correctamente")
	}

	// Esperar otro segundo (total 2 segundos)
	time.Sleep(1 * time.Second)

	// Ahora el primer request salió de la ventana, este debe ser permitido
	if !rl.IsAllowed("user2") {
		fmt.Println("❌ Request debería ser permitido después de la ventana")
	} else {
		fmt.Println("✓ Request permitido después de ventana deslizante")
	}
}

func testMultipleUsers() {
	rl := NewRateLimiter(2, 5*time.Second)

	// user3 hace 2 requests
	rl.IsAllowed("user3")
	rl.IsAllowed("user3")

	// user4 debe poder hacer requests independientemente
	if !rl.IsAllowed("user4") {
		fmt.Println("❌ user4 debería poder hacer requests")
	} else {
		fmt.Println("✓ Usuarios independientes funcionan correctamente")
	}

	// user3 no puede hacer más
	if rl.IsAllowed("user3") {
		fmt.Println("❌ user3 no debería poder hacer más requests")
	} else {
		fmt.Println("✓ Límite por usuario funciona correctamente")
	}
}

func testConcurrency() {
	rl := NewRateLimiter(10, 5*time.Second)
	var wg sync.WaitGroup
	successCount := 0
	var mu sync.Mutex

	// 20 goroutines intentando hacer requests
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if rl.IsAllowed("concurrent_user") {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	if successCount == 10 {
		fmt.Printf("✓ Concurrencia: %d/%d requests permitidos correctamente\n", successCount, 20)
	} else {
		fmt.Printf("❌ Concurrencia: %d requests permitidos, esperados 10\n", successCount)
	}
}
