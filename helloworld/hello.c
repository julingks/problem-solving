#include<stdio.h>
int main() {
    int cases;
    char name[1024];
    scanf("%d", &cases);
    while(cases--) {
        scanf("%s", name);
        printf("Hello, %s!\n", name);
    }
}
