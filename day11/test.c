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
    return old*19;
}
int op1(int old){
    return old+6;
}
int op2(int old){
    return old*old;
}
int op3(int old){
    return old+3;
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
    n0->value =  79;
    append(n0, 98);
    node_print(n0);
    struct Monkey m0 = {
        .list = n0,
        .op = &op0,
        .test = 23,
        .iftrue = 2,
        .iffals = 3,
    };

    node* n1 = newNode();
    n1->value = 54;
    append(n1, 65);
    append(n1, 75);
    append(n1, 74);
    struct Monkey m1 = {
        .list= n1,
        .op = &op1,
        .test = 19,
        .iftrue = 2,
        .iffals = 0,
    };

    node* n2 = newNode();
    n2->value = 79;
    append(n2, 60);
    append(n2, 97);
    struct Monkey m2 = {
        .list = n2,
        .op = &op2,
        .test = 13,
        .iftrue = 1,
        .iffals = 3,
    };


    node* n3 = newNode();
    n3->value = 74;
    struct Monkey m3 = {
        .list = n3,
        .op = &op3,
        .test = 17,
        .iftrue = 0,
        .iffals = 1,
    };

    struct Monkey* monkeys[4] = {&m0,&m1,&m2,&m3};
    int mc[4] = {0,0,0,0};

    for(int round = 0; round < 10000; round++){

        for(int i = 0; i < 4; i++) {


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
                //worry = worry / 3;
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

    printf("%d %d %d %d \n", mc[0], mc[1], mc[2],mc[3]);
}
