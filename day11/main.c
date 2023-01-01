#include <stdio.h>
#include <stdlib.h>

#define max_size 40

/*
 * for day11, I was bored with Go
 * wanted some change, so C it is.
 */
typedef int (*operation) (int old);


int op0(int old){
    printf("Worry level is multiplied by %d to %d.\n", old,old*19);
    return old*3;
}
int op1(int old){
    return old+8;
}
int op2(int old){
    return old+2;
}
int op3(int old){
    return old+4;
}
int op4(int old){
    return old*19;
}
int op5(int old){
    return old+5;
}
int op6(int old){
    return old*old;
}
int op7(int old){
    return old+1;
}

/* simple linked list */
struct node{
    struct node *next;
    int value;
};

typedef struct node node;

int len(node* n){
    if (n == NULL) {
        return 0; 
    }

    return 1 + len(n->next); 
}

node* newNode() {
    node *n = (node*) malloc(sizeof(node)); 
    n->value = 0;
    n->next = NULL;
    return n;
}

/* appends to the end */
void append(node* n, int i){
    if (n->next == NULL){
        n->next = (node*) malloc(sizeof(node)); 
        n->next->next = NULL;
        n->next->value = i;

        return;
    } 

    append(n->next,i);
}

void node_print(node *n){
    if (n == NULL) {
        printf("\n--\n"); 
        return;
    }
    printf(" => {%d} ", n->value);
    node_print(n->next);
}

struct Monkey{
    node *list;
    operation op;
    int test;
    int iftrue;
    int iffals;
};

int main() {

    node* n0 = newNode();
    n0->value =  65;
    append(n0, 78);
    node_print(n0);
    struct Monkey m0 = {
        .list = n0,
        .op = &op0,
        .test = 5,
        .iftrue = 2,
        .iffals = 3,
    };

    node* n1 = newNode();
    n1->value = 54;
    append(n1, 78);
    append(n1, 86);
    append(n1, 79);
    append(n1, 73);
    append(n1, 64);
    append(n1, 85);
    append(n1, 88);
    struct Monkey m1 = {
        .list= n1,
        .op = &op1,
        .test = 11,
        .iftrue = 4,
        .iffals = 7,
    };

    node* n2 = newNode();
    n2->value = 69;
    append(n2, 97);
    append(n2, 77);
    append(n2, 88);
    append(n2, 87);
    struct Monkey m2 = {
        .list = n2,
        .op = &op2,
        .test = 2,
        .iftrue = 5,
        .iffals = 3,
    };


    node* n3 = newNode();
    n3->value = 99;
    struct Monkey m3 = {
        .list = n3,
        .op = &op3,
        .test = 13,
        .iftrue = 1,
        .iffals = 5,
    };

    node* n4 = newNode();
    n4->value = 60;
    append(n4, 57);
    append(n4, 52);
    struct Monkey m4 = {
        .list = n4,
        .op = &op4,
        .test = 7,
        .iftrue = 7,
        .iffals = 6,
    };

    node* n5 = newNode();
    n5->value = 91;
    append(n5, 82);
    append(n5, 85);
    append(n5, 73);
    append(n5, 84);
    append(n5, 53);
    struct Monkey m5 = {
        .list = n5,
        .op = &op5,
        .test = 3,
        .iftrue = 4,
        .iffals = 1,
    };


    node* n6 = newNode();
    n6->value = 88;
    append(n6, 74);
    append(n6, 68);
    append(n6, 56);
    struct Monkey m6 = {
        .list = n6,
        .op = &op6,
        .test = 17,
        .iftrue = 0,
        .iffals = 2,
    };


    node* n7 = newNode();
    n7->value = 54;
    append(n7, 82);
    append(n7, 72);
    append(n7, 71);
    append(n7, 53);
    append(n7, 99);
    append(n7, 67);
    struct Monkey m7 = {
        .list = n7,
        .op = &op7,
        .test = 19,
        .iftrue = 6,
        .iffals = 0,
    };

    struct Monkey* monkeys[8] = {&m0,&m1,&m2,&m3,&m4,&m5,&m6,&m7};
    int mc[8] = {0,0,0,0,0,0,0,0};

    for(int round = 0; round < 20; round++){

        for(int i = 0; i < 8; i++) {


            /*
             * first round
             * in each step, divide worry level by 3
             */

            struct Monkey* m = monkeys[i];
            printf("Monkey %d:\n", i);
            /* loop until a monkey has no remaining item */

            while(m->list != NULL ) {
                mc[i]++;
                printf("Monkey inspects an item with a worry level of %d:\n",m->list->value);
                int worry = m->op(m->list->value);
                printf("Monkey gets bored with the item-> Worry level is divided by 3");
                worry = worry / 3;
                printf(" to %d \n", worry);
                if( (worry % m->test) == 0 ) {
                    printf("item worry level %d is divisable by %d\n", worry, m->test);
                    struct Monkey* nm = monkeys[m->iftrue];
                    /* delete from the current monkey, then add to next monkey */
                    m->list = m->list->next;
                    if(nm->list == NULL){
                        nm->list = newNode();
                        nm->list->value = worry;
                    }else{
                        append(nm->list, worry);
                    }
                    printf("item with worry level %d is thrown to Monkey %d.\n", worry, m->iftrue);
                    printf("new monkeys list: ---\n");
                } else {
                    printf("item worry level %d is not divisable by %d\n", worry, m->test);
                    struct Monkey* nm = monkeys[m->iffals];
                    /* delete from the current monkey, then add to next monkey */
                    m->list = m->list->next;
                    if(nm->list == NULL){
                        nm->list = newNode();
                        nm->list->value = worry;
                    }else{
                        append(nm->list, worry);
                    }
                    printf("item with worry level %d is thrown to Monkey %d.\n", worry, m->iffals);
                    printf("new monkeys list: ---\n");
                }

            }


        }
        printf("After round %d, the monkeys are holding items with these worry levels:\n", round);
        node_print(m0.list); 
        node_print(m1.list); 
        node_print(m2.list); 
        node_print(m3.list); 
    }

    printf("%d %d %d %d %d %d %d %d \n",
            mc[0], mc[1], mc[2],mc[3],mc[4],mc[5],mc[6],mc[7]);
}
