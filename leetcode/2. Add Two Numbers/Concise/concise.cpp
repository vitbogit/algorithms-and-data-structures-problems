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
        ListNode* l3 = new ListNode();
        ListNode* head = l3;
        for(int carry=0; l1 != NULL || l2 != NULL || carry > 0; carry /= 10){
            if (l1 != NULL){
                carry += l1->val;
                l1 = l1->next;
            }
            if (l2 != NULL){
                carry += l2->val;
                l2 = l2->next;
            }
            l3->next = new ListNode(carry%10);
            l3 = l3->next;
        }
        return head->next;
    }
};