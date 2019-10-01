#include <stdio.h>
#include <stdlib.h>

int main( ) {
  printf("===========================\n");
  printf("Welcome to the Curio Shop!\n");
  printf("===========================\n");
  int selection = 1;
  int money = 100;
  while (selection != 4) {
    printf("\n");
    if (selection == 1) {
      printf("You have $%d to spend.\n", money);
      printf("[1] Print this menu\n");
      printf("[2] List shop items\n");
      printf("[3] Buy something\n");
      printf("[4] Exit\n");
    } else if (selection == 2) {
      printf("We have a few items in stock today!\n");
      printf("Item 1: Cyclone Frisbee, $10\n");
      printf("Item 2: PG Tuition, $22,460\n");
      printf("Item 3: Flag, $1,000,000\n");
      printf("Enter [3] to buy something!\n");
    } else if (selection == 3) {
      printf("Which item do you want to buy?\n");
      int item;
      scanf("%d", &item);
      printf("How many do you want to buy?\n");
      int quantity;
      scanf("%d", &quantity);
      int price;
      if (quantity < 0) {
        printf("%d is invalid! You'll have to try harder than that...\n", &quantity);
        continue;
      }
      if (item == 1) {
        price = quantity * 10;
      } else if (item == 2) {
        price = quantity * 22460;
      } else if (item == 3) {
        price = quantity * 1000000;
      } else {
        printf("That is not a valid item.\n");
        continue;
      }
      if (price > money) {
        printf("You don't have enough money for that (yet).\n");
        continue;
      }
      money -= price;
      printf("Payment received!\n");
      if (item == 3) {
        printf("Your flag is: pgctf{2147483648_is_a_mag1c_numb3r}\n");
        exit(0);
      } else {
        printf("Your new balance is %d\n", money);
      }
    } else if (selection == 4) {
      exit(0);
    } else {
      printf("That is not a valid option.\n");
    }
    scanf("%d", &selection);
  }
  return 0;
}