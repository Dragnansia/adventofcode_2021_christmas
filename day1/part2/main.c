#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void insert_last_val(int *windows, int new_val) {
    int tmp1 = windows[1], tmp2 = windows[2];
    windows[0] = tmp1;
    windows[1] = tmp2;
    windows[2] = new_val;
}

int add_all(int *arr, int len) {
    int res = 0;
    for (int i = 0; i < len; ++i) res += arr[i];
    return res;
}

int main(int argc, char **argv) {
    FILE *values = fopen("./values", "r");
    if (values == NULL) {
        printf("Can't open ./values file");
        exit(EXIT_FAILURE);
    }

    int c_depth_increase = 0, previous_value = INT_MAX, windows[3] = {0};

    char *line = NULL;
    size_t len = 0;
    ssize_t read;

    for (int i = 0; i < 2; ++i) {
        read = getline(&line, &len, values);
        insert_last_val(windows, strtol(line, NULL, 10));
    }

    while ((read = getline(&line, &len, values)) != -1) {
        insert_last_val(windows, strtol(line, NULL, 10));
        int val = add_all(windows, 3);

        if (val > previous_value) {
            c_depth_increase += 1;
        }

        previous_value = val;
    }

    printf("Depth Increases: %d\n", c_depth_increase);
    return 0;
}
