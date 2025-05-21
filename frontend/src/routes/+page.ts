import { GetAdmins, GetStudents } from '$lib/wailsjs/go/main/App';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	const admins = await GetAdmins();
	const students = await GetStudents();
	return { admins, students };
};
