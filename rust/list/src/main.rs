

// Definition for singly-linked list.
#[derive(Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}
// 
impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

struct Solution();
impl Solution {
    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
		let mut cur = head?;
		if let Some(mut next) = cur.next {
			cur.next = Solution::swap_pairs(next.next);
			next.next = Some(cur);
			return Some(next)
		}
		Some(cur)
    }


	pub fn reverse_between(head: Option<Box<ListNode>>, left: i32, right: i32) -> Option<Box<ListNode>> {

		let mut head = head;
		let mut left_prev_node = &mut head;
		for _ in 0..left-2{
			left_prev_node = &mut left_prev_node.as_mut().unwrap().next;
		}
		let mut left_prev_node_ptr = left_prev_node.as_mut().unwrap();
		// if let Some(left_node_ptr) = left_node {
			left_prev_node_ptr.next = Solution::reverse(left_prev_node_ptr.next.take());
		// }

		head

    }



	pub fn _reverse(prev: Option<Box<ListNode>>, cur: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
		let mut cur_node = cur?;
		let next = cur_node.next;
		cur_node.next = prev;
		if next.is_none() {
			return Some(cur_node);
		}
		Solution::_reverse(Some(cur_node), next)
	}

	pub fn reverse(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
		Solution::_reverse(None, head)
	}


	pub fn print_list(head: &Option<Box<ListNode>>) {
		let mut cur = head;
		while let Some(cur_ptr) = cur {
			print!("{} ",cur_ptr.val);
			cur = &cur_ptr.next;
		}
		println!()
	}
}

fn main() {
	// let list = None;
	let list = Some(Box::new(ListNode{
		val: 1,
		// next: None,
		next: Some(Box::new(ListNode{
			val:2,
			next: Some(Box::new(ListNode{
				val:3,
				// next: None,
				next:Some(Box::new(ListNode{
					val:4,
					next: Some(Box::new(ListNode{
						val: 5,
						next: None
					}))
				}))
			}))
		}))
	}));

	Solution::print_list(&list);
	let new_list = Solution::reverse_between(list, 2,4);
	Solution::print_list(&new_list);



}
