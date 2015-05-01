#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    char* format;
    unsigned int length;
    unsigned char* bytes;
} image;

image* process(image*);

#ifdef __cplusplus
}
#endif