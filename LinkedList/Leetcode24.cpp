#include <iostream>

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
    while (temp) {
      temp = temp->next;
    }
    temp->next = newNode;
    newNode->next = nullptr;

    return head;
  }

  void displayLL() {

    if (!head) {
      std::cout << "List is empty" << std::endl;
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

class Leetcode24 {};
