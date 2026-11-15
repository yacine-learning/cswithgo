package main

// Plan groups topics into semesters.
//
// deps[t] lists the prerequisites of topic t: each of them must be in an earlier semester
// than t. A topic that appears only as a prerequisite is a topic too.
//
//   - semester 0 holds every topic with no prerequisite; semester k holds every topic whose
//     prerequisites are all in semesters before k (each topic as early as possible);
//   - the topics of each semester are sorted by name;
//   - when some topics can never be studied (they are on a cycle, or need one),
//     Plan returns the error "cannot order: a, b, c": all those topics, sorted.
func Plan(deps map[string][]string) ([][]string, error) {
	return nil, nil
}
