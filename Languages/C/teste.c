#include <stdio.h>

int main() {    
    float lenght;
    float height;
    printf("What\'s the rectangle\'s lenght?\n");
    scanf("%f", &lenght);
    printf("What\'s the rectangle\'s height?\n");
    scanf("%f", &height);        
    printf("The perimeter of the rectangle is %0.2f\n", lenght * 2 + height * 2 );
    return 0;
}

