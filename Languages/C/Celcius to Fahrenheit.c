#include <stdio.h>

int main() {
    float Celcius;

    printf("What is the degree in Celcius??\n");
    scanf("%f", &Celcius);
    
    Celcius = (Celcius * 1.8 + 32);
    printf("The degree is %.2f Fahrenheit!!\n", Celcius);
}
