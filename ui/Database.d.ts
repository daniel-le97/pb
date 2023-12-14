/**
 * This file was @generated using typed-pocketbase
 */

// https://pocketbase.io/docs/collections/#base-collection
interface BaseCollectionRecord {
  id: string
  created: string
  updated: string
}

// https://pocketbase.io/docs/collections/#auth-collection
interface AuthCollectionRecord {
  id: string
  created: string
  updated: string
  username: string
  email: string
  emailVisibility: boolean
  verified: boolean
}

// https://pocketbase.io/docs/collections/#view-collection
interface ViewCollectionRecord {
  id: string
}

// utilities

type MaybeArray<T> = T | T[]

// ===== users =====

export type UsersResponse = {
  name?: string
  avatar?: string
} & AuthCollectionRecord

export interface UsersCreate {
  name?: string
  avatar?: string
}

export interface UsersUpdate {
  name?: string
  avatar?: string
}

export interface UsersCollection {
  type: 'auth'
  collectionId: '_pb_users_auth_'
  collectionName: 'users'
  response: UsersResponse
  create: UsersCreate
  update: UsersUpdate
  relations: {
    'teams(users)': TeamsCollection[]
  }
}

// ===== teams =====

export type TeamsResponse = {
  users?: Array<string>
} & BaseCollectionRecord

export interface TeamsCreate {
  users?: MaybeArray<string>
}

export interface TeamsUpdate {
  users?: MaybeArray<string>
  'users+'?: MaybeArray<string>
  'users-'?: MaybeArray<string>
}

export interface TeamsCollection {
  type: 'base'
  collectionId: 'qtgju6qow8zu5sc'
  collectionName: 'teams'
  response: TeamsResponse
  create: TeamsCreate
  update: TeamsUpdate
  relations: {
    users: UsersCollection[]
  }
}

// ===== projects =====

export type ProjectsResponse = {
  repoURL?: string
  name?: string
  deployed?: boolean
  buildpack?: 'nixpacks' | 'dockerfile' | 'docker-compose'
  configured?: boolean
  baseDir?: string
  buildDir?: string
  https?: boolean
  www?: boolean
  managed?: boolean
  installCommand?: string
  buildCommand?: string
  startCommand?: string
  ports?: string
  exposedPorts?: string
} & BaseCollectionRecord

export interface ProjectsCreate {
  repoURL?: string | URL
  name?: string
  deployed?: boolean
  buildpack?: 'nixpacks' | 'dockerfile' | 'docker-compose'
  configured?: boolean
  baseDir?: string
  buildDir?: string
  https?: boolean
  www?: boolean
  managed?: boolean
  installCommand?: string
  buildCommand?: string
  startCommand?: string
  ports?: string
  exposedPorts?: string
}

export interface ProjectsUpdate {
  repoURL?: string | URL
  name?: string
  deployed?: boolean
  buildpack?: 'nixpacks' | 'dockerfile' | 'docker-compose'
  configured?: boolean
  baseDir?: string
  buildDir?: string
  https?: boolean
  www?: boolean
  managed?: boolean
  installCommand?: string
  buildCommand?: string
  startCommand?: string
  ports?: string
  exposedPorts?: string
}

export interface ProjectsCollection {
  type: 'base'
  collectionId: 'tqlcamhhas2xzr7'
  collectionName: 'projects'
  response: ProjectsResponse
  create: ProjectsCreate
  update: ProjectsUpdate
  relations: {
    'logs(project)': LogsCollection[]
    'queue(project)': QueueCollection[]
  }
}

// ===== logs =====

export type LogsResponse = {
  project: string
  content?: string
  buildTime?: number
} & BaseCollectionRecord

export interface LogsCreate {
  project: string
  content?: string
  buildTime?: number
}

export interface LogsUpdate {
  project?: string
  content?: string
  buildTime?: number
  'buildTime+'?: number
  'buildTime-'?: number
}

export interface LogsCollection {
  type: 'base'
  collectionId: '6po51li0eihk83d'
  collectionName: 'logs'
  response: LogsResponse
  create: LogsCreate
  update: LogsUpdate
  relations: {
    project: ProjectsCollection
  }
}

// ===== templates =====

export type TemplatesResponse = {
  type?: 'portainer' | 'caprover' | 'custom'
  content?: string
} & BaseCollectionRecord

export interface TemplatesCreate {
  type?: 'portainer' | 'caprover' | 'custom'
  content?: string
}

export interface TemplatesUpdate {
  type?: 'portainer' | 'caprover' | 'custom'
  content?: string
}

export interface TemplatesCollection {
  type: 'base'
  collectionId: '67nwlz2znzqld2l'
  collectionName: 'templates'
  response: TemplatesResponse
  create: TemplatesCreate
  update: TemplatesUpdate
  relations: {}
}

// ===== git =====

export type GitResponse = {
  app_id?: string
  webhook_secret?: string
  private_key?: string
} & BaseCollectionRecord

export interface GitCreate {
  app_id?: string
  webhook_secret?: string
  private_key?: string
}

export interface GitUpdate {
  app_id?: string
  webhook_secret?: string
  private_key?: string
}

export interface GitCollection {
  type: 'base'
  collectionId: '5c1oykad2csm1dk'
  collectionName: 'git'
  response: GitResponse
  create: GitCreate
  update: GitUpdate
  relations: {}
}

// ===== queue =====

export type QueueResponse = {
  project?: string
  active?: boolean
} & BaseCollectionRecord

export interface QueueCreate {
  project?: string
  active?: boolean
}

export interface QueueUpdate {
  project?: string
  active?: boolean
}

export interface QueueCollection {
  type: 'base'
  collectionId: 'ojmqge2b4cj4ywj'
  collectionName: 'queue'
  response: QueueResponse
  create: QueueCreate
  update: QueueUpdate
  relations: {
    project: ProjectsCollection
  }
}

// ===== Schema =====

export interface Schema {
  users: UsersCollection
  teams: TeamsCollection
  projects: ProjectsCollection
  logs: LogsCollection
  templates: TemplatesCollection
  git: GitCollection
  queue: QueueCollection
}
