#include <iostream>
#include <ostream>

struct Node {
  int data;
  Node *next;
};

class LinkedList {
  Node *head;

public:
  LinkedList() : head(nullptr) {};

  Node *insertAtEnd(int val) {
    Node *newNode = new Node();
    newNode->data = val;
    newNode->next = nullptr;

    if (!head) {
      head = newNode;
      return head;
    }
    Node *temp = head;
    while (temp->next) {
      temp = temp->next;
    }
    temp->next = newNode;
    newNode->next = nullptr;
    return head;
  }

  void displayLL() {
    if (!head) {
      std::cout << "List is empty";
    }
    Node *temp = head;
    while (temp) {
      std::cout << temp->data << "->";
      temp = temp->next;
    }
    std::cout << "nullptr" << std::endl;
  }

  Node *getHead() { return head; }
};

class Leetcode86 {

public:
  static Node *partition(Node *head, int x) {

    Node *dummy = new Node();
    dummy->data = -1;
    dummy->next = nullptr;
    Node *temp = head;
    Node *traverser = dummy;

    while (temp) {

      if (temp->data < x) {
        traverser->next = temp;
        traverser = traverser->next;
        traverser->next = nullptr;
      }
      temp = temp->next;
    }

    temp = head;
    while (temp) {

      if (temp->data >= x) {
        traverser->next = temp;
        traverser = traverser->next;
        traverser->next = nullptr;
      }

      temp = temp->next;
    }

    return dummy->next;
  }
};

int main() {

  LinkedList list;

  list.insertAtEnd(1);
  list.insertAtEnd(4);
  list.insertAtEnd(3);
  list.insertAtEnd(2);
  list.insertAtEnd(5);
  list.insertAtEnd(2);

  std::cout << "List 1: ";
  list.displayLL();

  Node *result = Leetcode86 ::partition(list.getHead(), 3);

  LinkedList resultList;
  Node *temp = result;

  while (temp) {
    resultList.insertAtEnd(temp->data);
    temp = temp->next;
  }
  resultList.displayLL();
}
