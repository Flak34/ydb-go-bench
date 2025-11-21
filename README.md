# YDB Concurrent Result Sets Benchmark

Бенчмарк сравнивает три режима выполнения запросов в `ydb-go-sdk/v3`:

1. **Один большой запрос + ConcurrentResultSets = true**  
2. **Один большой запрос + ConcurrentResultSets = false**  
3. **Серия последовательных запросов вместо одного мультиселекта**

Цель — измерить влияние параллельной обработки ResultSet'ов на производительность.

---

## 📁 Структура

- **`main_test.go`** — инициализация драйвера и генерация набора SELECT-запросов.  
- **`concurrent_result_sets_test.go`** — сами бенчмарки.

---

## ▶️ Запуск бенчмарков

Обычный запуск:

```sh
go test -bench=. -benchmem
