import { GraphQLStudent } from './model';
import { studentList } from './list';
import { studentQuery } from './query';
import { studentMutation } from './mutation';

const Student = GraphQLStudent;
const studentLists = studentList;
const studentQueries = studentQuery;
const studentMutations = studentMutation;

export { Student, studentLists, studentQueries, studentMutations };
