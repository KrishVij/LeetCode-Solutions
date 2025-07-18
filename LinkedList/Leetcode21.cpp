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
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode* dummy = new ListNode(-1);
        ListNode* temp1 = list1;
        ListNode* temp2 = list2;
        ListNode* temp3 = dummy;
        while (temp1 != nullptr && temp2 != nullptr) {
            if (temp1->val < temp2->val) {
                temp3->next = temp1;
                temp1 = temp1->next;
                temp3 = temp3->next;
            } else {
                temp3->next = temp2;
                temp2 = temp2->next;
                temp3 = temp3->next;
            }
        }
        if (temp1 != nullptr) {
            temp3->next = temp1;
        }
        if (temp2 != nullptr) {
            temp3->next = temp2;
        }
        return dummy->next;
    }
    
};
