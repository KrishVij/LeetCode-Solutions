#include<queue>
#include<iostream>

struct Node{
    int data;
    Node* next;
    Node* child;
};

class multiLL{
    Node* head;

public :
    multiLL() : head(nullptr){};

    Node* insertAtNext(int val){
        Node* newNode = new Node ();
        newNode -> data = val;
        if (!head){
            head = newNode;
            return head;
        }
        Node* temp = head;
        while (temp -> next){
            temp = temp -> next;
        }
        temp -> next = newNode;
        newNode -> next = nullptr;
        return head;

    }
    Node* insertAtChild(Node* parent,int val){
        if (!parent){
        std::cout<<"Parent doesn't exist"<<std::endl;
        }
        Node* newNode = new Node ();
        newNode -> data = val;
        if (!parent -> child){
            parent -> child = newNode;
            return head;
        }
        Node* holder = parent -> child;
        while (holder -> child){
            holder = holder -> child;
        }
        holder -> child = newNode;
        newNode -> child = nullptr;
        return head;
    }

    void displayIteratively(){
        if (!head){
            std::cout<<"List is empty";
        }
        std::queue<Node*> storage;
        storage.push(head);
        while (!storage.empty()){
            Node* current = storage.front();
            storage.pop();
            std::cout<<current -> data<<" ";
            if (current -> child){
                storage.push(current -> child);
            }
            if (current -> next){storage.push(current -> next);}
        }
        std::cout<<std::endl;
    }
   Node* getHead(){return head;} 

};
class flatten{

public :

    Node* merge(Node* list1,Node* list2){
        Node* dummy = new Node ();
        dummy -> data = -1;
        dummy -> next = nullptr;
        dummy -> child = nullptr;
        Node* temp = dummy;
        
        while (list1 && list2){
            if (list1 -> data < list2 -> data){
                temp -> child = list1;
                temp = list1;
                list1 = list1 -> next;
            }else{
                temp -> child = list2;
                temp = list2;
                list2 = list2 -> next;
            }
            temp -> next = nullptr;
        }
        if (list1){temp -> child = list1;}
        if (list2){temp -> child = list2;}
        
        return dummy -> child;
    }
    Node* flattenLinkedList(Node* head){
        if (!head || !head -> next){
            return head;
        }
        Node* mergeHead = flattenLinkedList(head -> next);
       head =  merge(head,mergeHead);
        return head;
    }
};
int main(){
    multiLL list;

    list.insertAtNext(48);
    list.insertAtNext(39);
    list.insertAtNext(41);

    list.insertAtChild(list.getHead(),40);
    Node* secondHead = list.getHead() -> next;
    list.insertAtChild(secondHead,60);
    list.insertAtChild(secondHead, 55); 

    list.displayIteratively();

    flatten flattener;
    Node* flattenedHead = flattener.flattenLinkedList(list.getHead());

    std::cout << "\nFlattened Linked List:\n";
    Node* temp = flattenedHead;
    while (temp) {
        std::cout << temp->data << " ";
        temp = temp->child;
    }
    std::cout << std::endl;

    return 0;

}

