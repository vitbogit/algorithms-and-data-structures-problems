# Concise solution

It\'s pretty straightforward. Just make sure you take two things into account:
- a concise solution is better than an intuitive one (creating integer variables representing given numbers), since in some tests the numbers are too large even for int64;
- don't forget to add the remaining amount to l3, even if l1 and l2 are already nil;
- we need to somehow deal with the “empty” node at the beginning of created l3.
