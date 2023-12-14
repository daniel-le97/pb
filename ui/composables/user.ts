import type { ProjectsResponse } from '../Database'

export const useUser = () => useState('user')

export const useActiveProject = () => useState<ProjectsResponse>('activeProject')
