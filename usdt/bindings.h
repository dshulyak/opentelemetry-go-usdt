#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

void stacks_enter(uint64_t span_id,
                  uint64_t parent_span_id,
                  uint64_t id,
                  uint64_t amount,
                  const char *span_name);

void stacks_exit(uint64_t span_id);
