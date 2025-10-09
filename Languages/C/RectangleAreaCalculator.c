#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(){ 
    char question[10];
    char area[] = "area";
    char perimeter[] = "perimeter";
    int height;
    int lenght;

    printf("Do you want to calculate the area or perimeter of the rectangle?\n");
    scanf("%s", &question);
    //strcmp(question,area) == 

    printf("What's the height of the rectangle?\n");
    scanf("%d", &height);
    printf("And what is the lenght of the rectangle?\n");
    scanf("%d", &lenght);

    if (strcmp(question,area) == 0) {
        printf("The area of the rectangle is %d\n", height * lenght);
        return 0;
    } else if (strcmp(question,perimeter) == 0) {
          printf("The perimeter of the rectangle is %d\n", height * 2 + lenght * 2 );
          return 0;
    } else {
          printf("Invalid response, please wright either \'area\' or \'perimeter\'\n");
          return 0;
    }

}
