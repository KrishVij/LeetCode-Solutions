/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* temp1 = l1;
        ListNode* temp2 = l2;
        ListNode* head = nullptr;
        ListNode* temp3 = nullptr;
        int carry{0};
        int val1{0};
        int val2{0};

        while (temp1 != nullptr || temp2 != nullptr || carry > 0) {
            if (temp1 != nullptr) {
                val1 = temp1->val;
            } else {
                val1 = 0;
            }
            if (temp2 != nullptr) {
                val2 = temp2->val;
            } else {
                val2 = 0;
            }

            int sum = val1 + val2 + carry;
            carry = sum / 10;
            int digit = sum % 10;
            ListNode* newNode = new ListNode(digit);

            if (head == nullptr) {
                head = newNode;
                temp3 = head;
            } else {
                temp3->next = newNode;
                temp3 = temp3->next;
            }
            if (temp1 != nullptr)
                temp1 = temp1->next;
            if (temp2 != nullptr)
                temp2 = temp2->next;
        }
        return head;
    }
};
