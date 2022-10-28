import { GraphQLStudent } from './model';
import { studentConnection } from './list';
import { StudentQueries } from './query';
import { studentMutation } from './mutation';

// exporting graphql model of student.
export const Student = GraphQLStudent;

// exporting student connection
export const StudentConnection = studentConnection;

// exporting student queries
export const studentQueries = StudentQueries;

// exporting mutations of student
export const studentMutations = studentMutation;
