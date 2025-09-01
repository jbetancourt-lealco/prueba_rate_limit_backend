# Backend de Rate Limiting - Implementación en Go

## Descripción General

Este es un proyecto de evaluación técnica que implementa un **Sistema de Rate Limiting con Ventana Deslizante** en Go. El sistema controla el número de requests por usuario dentro de una ventana de tiempo configurable, proporcionando operaciones thread-safe y optimización de memoria.

## Enunciado del Problema

Implementar un sistema de rate limiting que controle el número de requests por usuario usando un enfoque de ventana de tiempo deslizante.

### Requisitos Principales

- ✅ Limitar requests a N por usuario en los últimos M segundos
- ✅ `IsAllowed(userID string) bool` - Verificar si el request está permitido
- ✅ `Reset(userID string)` - Limpiar el historial de requests del usuario
- ✅ Operaciones thread-safe (manejar concurrencia)
- ✅ Optimización de memoria (limpiar datos antiguos)

### Características Adicionales

- 🔄 `GetStats()` - Estadísticas y métricas del usuario
- 🔄 Diferentes límites por tipo de usuario (gratuito/premium)
- 🔄 Limpieza de usuarios inactivos

## Arquitectura

### Estructuras de Datos

El sistema utiliza una estructura `RateLimiter` con:
- `maxRequests`: Máximo de requests permitidos por ventana
- `windowDuration`: Tamaño de la ventana de tiempo
- `mu`: Mutex para seguridad en concurrencia
- Seguimiento del historial de requests del usuario (por implementar)

### Métodos Principales

```go
type RateLimiter struct {
    maxRequests    int
    windowDuration time.Duration
    mu            sync.Mutex
}

// Funcionalidad principal
func (rl *RateLimiter) IsAllowed(userID string) bool
func (rl *RateLimiter) Reset(userID string)

// Características adicionales
func (rl *RateLimiter) GetStats(userID string) map[string]interface{}
func (rl *RateLimiter) CleanupInactiveUsers(inactivityThreshold time.Duration) int
```

## Ejemplo de Uso

```go
// Crear rate limiter: 5 requests por 10 segundos
rl := NewRateLimiter(5, 10*time.Second)

// Verificar si el usuario puede hacer un request
if rl.IsAllowed("user123") {
    // Procesar request
    fmt.Println("Request permitido")
} else {
    // Límite de rate excedido
    fmt.Println("Límite de rate excedido")
}

// Resetear historial del usuario
rl.Reset("user123")
```

## Escenarios de Prueba

El proyecto incluye casos de prueba exhaustivos:

### 1. Funcionamiento Básico
- Verificar que los primeros N requests sean permitidos
- Verificar que el request N+1 sea bloqueado
- Probar con 3 requests en ventana de 5 segundos

### 2. Ventana Deslizante
- Probar expiración basada en tiempo
- Verificar que los requests sean bloqueados dentro de la ventana
- Verificar que los requests sean permitidos después de que expire la ventana

### 3. Múltiples Usuarios
- Asegurar aislamiento entre usuarios
- Probar límites de rate independientes por usuario
- Verificar que no haya interferencia entre usuarios

### 4. Concurrencia
- Probar seguridad de threads con 20 goroutines
- Verificar que exactamente N requests sean permitidos
- Prueba de estrés con acceso concurrente

## Ejecutar el Proyecto

### Prerrequisitos
- Go 1.19+ (usa características modernas de Go)

### Ejecución
```bash
# Navegar al directorio del proyecto
cd prueba_rate_limit_backend

# Ejecutar la aplicación
go run main.go
```

### Salida Esperada
```
=== Test 1: Funcionamiento Básico ===
✓ Request 1 permitido
✓ Request 2 permitido
✓ Request 3 permitido
✓ Request 4 bloqueado correctamente

=== Test 2: Ventana Deslizante ===
✓ Request bloqueado correctamente
✓ Request permitido después de ventana deslizante

=== Test 3: Múltiples Usuarios ===
✓ Usuarios independientes funcionan correctamente
✓ Límite por usuario funciona correctamente

=== Test 4: Concurrencia ===
✓ Concurrencia: 10/20 requests permitidos correctamente
```

## Estado de Implementación

### ✅ Completado
- Estructura del proyecto e interfaz
- Framework de pruebas y escenarios
- Firmas de métodos y documentación

### 🔄 Por Implementar
- Almacenamiento del historial de requests del usuario
- Lógica de ventana deslizante en `IsAllowed()`
- Operaciones thread-safe
- Mecanismos de limpieza de memoria
- Recolección de estadísticas
- Diferenciación por tipo de usuario

## Consideraciones Técnicas

### Gestión de Memoria
- Implementar estructuras de datos eficientes para almacenar timestamps
- Limpieza regular de entradas expiradas
- Considerar usar buffers circulares o mapas basados en tiempo

### Concurrencia
- Usar mutex para seguridad de threads
- Considerar locks de lectura-escritura para mejor rendimiento
- Manejar condiciones de carrera en estadísticas

### Rendimiento
- Optimizar comparaciones de timestamps
- Minimizar asignaciones de memoria
- Algoritmos de limpieza eficientes

## Mejoras Futuras

1. **Integración con Redis**: Almacenar datos de rate limit en Redis para sistemas distribuidos
2. **Configuración**: Cargar límites desde archivos de configuración o variables de entorno
3. **Métricas**: Integración con Prometheus para monitoreo
4. **Servidor API**: Endpoints HTTP para verificaciones de rate limiting
5. **Middleware**: Middleware HTTP de Go para fácil integración

## Estructura del Proyecto

```
prueba_rate_limit_backend/
├── main.go          # Implementación principal y pruebas
├── go.mod           # Definición del módulo Go
└── README.md        # Este archivo
```

## Contribución

Este es un proyecto de evaluación técnica. La implementación se enfoca en:
- Código limpio y legible
- Cobertura exhaustiva de pruebas
- Manejo apropiado de errores
- Consideraciones de rendimiento
- Patrones listos para producción

## Licencia

Este proyecto es creado para propósitos de evaluación técnica.
