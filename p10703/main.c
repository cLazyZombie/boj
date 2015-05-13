#include <stdio.h>
#include <stdlib.h>

main(){
    int row, column;
    scanf("%d %d", &row, &column);

    // height fields
    int* meteors = (int*)malloc(sizeof(int) * column);
    int* grounds = (int*)(malloc(sizeof(int) * column));

    int x, y;

    for (x = 0; x < column; x++) {
        meteors[x] = -1;
        grounds[x] = row;
    }

    // store separately meteor and other
    char* mpic = (char*)malloc(sizeof(char)*row*column);
    char* opic = (char*)malloc(sizeof(char)*row*column);

    char* line = (char*)malloc(sizeof(char)*column+2);
    for (y = 0; y < row; y++) {
        scanf("%s", line);
        //fgets(line, column+2, stdin);

        for (x = 0; x < column; x++) {
            int index = y * column + x;
            char b = line[x];
            switch (b){
                case 'X':
                mpic[index] = b;
                opic[index] = '.';

                if (y > meteors[x]) {
                    meteors[x] = y;
                }

                break;

                case '#':
                opic[index] = b;
            	if (y < grounds[x]) {
            	       grounds[x] = y;
            	}
                break;

                case '.':
                opic[index] = b;
                break;
            }
        }
    }

    // get min distance
    int minDistance = row;
	for (x = 0; x < column; x++) {
        if (meteors[x] == -1){
            continue;
        }

        int d = grounds[x] - meteors[x] - 1;
        if (d < minDistance) {
            minDistance = d;
        }
    }

    // print result pic
    for (y = 0; y < row; y++) {
        for (x = 0; x < column; x++) {
            int offY = y - minDistance;
            if (offY >= 0 && mpic[offY*column+x] == 'X') {
                putchar('X');
            } else {
                putchar(opic[y*column+x]);
            }
        }
        putchar('\n');
    }
}
