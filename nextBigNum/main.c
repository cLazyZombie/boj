#include <stdio.h>
#include <stdlib.h>
#include <string.h>

main() {
    int n;
    scanf("%d", &n);
    printf("%d\n", nextSameDigits(n));
}

void sort(char* p, int len) {
    int i, j;
    for (i = 0; i < len -1; i++){
        for (j = i+1; j < len; j++){
            if (p[i] > p[j]) {
                char t = p[i];
                p[i] = p[j];
                p[j] = t;
            }
        }
    }
}

int nextSameDigits(int I) {
    char a[33];
    int len = strlen(a);

    int i, j;
    for (i = len-2; i >= 0; i--) {
        char c = a[i];
        char small = '9' + 1;
        int swapIndex = -1;
        for (j = i + 1; j < len; j++) {
            if (a[j] > c && a[j] < small) {
                small = a[j];
                swapIndex = j;
            }
        }

        if (swapIndex != -1) {
            char temp = a[i];
            a[i] = a[swapIndex];
            a[swapIndex] = temp;

            sort(&(a[i+1]), len - i - 1);
            return atoi(a);
        }
    }

    return -1;
}
