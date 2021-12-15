#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char **argv) {
    FILE *values = fopen("./values", "r");
    if (values == NULL) {
        printf("Can't open ./values file");
        exit(EXIT_FAILURE);
    }

    int c_depth_increase = 0, previous_value = 2147483647;

    char *line = NULL;
    size_t len = 0;
    ssize_t read;
    while ((read = getline(&line, &len, values)) != -1) {
        int val = strtol(line, NULL, 10);
        if (val > previous_value) {
            c_depth_increase += 1;
        }

        previous_value = val;
    }

    printf("Depth Increases: %d\n", c_depth_increase);
    return 0;
}
