/**
 * This file was @generated using typed-pocketbase
 */

// https://pocketbase.io/docs/collections/#base-collection
type BaseCollectionRecord = {
	id: string;
	created: string;
	updated: string;
};

// https://pocketbase.io/docs/collections/#auth-collection
type AuthCollectionRecord = {
	id: string;
	created: string;
	updated: string;
	username: string;
	email: string;
	emailVisibility: boolean;
	verified: boolean;
};

// https://pocketbase.io/docs/collections/#view-collection
type ViewCollectionRecord = {
	id: string;
};

// utilities

type MaybeArray<T> = T | T[];

// ===== users =====

export type UsersResponse = {
    name?: string;
	avatar?: string;
} & AuthCollectionRecord;

export type UsersCreate = {
	name?: string;
	avatar?: string;
};

export type UsersUpdate = {
	name?: string;
	avatar?: string;
};

export type UsersCollection = {
	type: 'auth';
	collectionId: '_pb_users_auth_';
	collectionName: 'users';
	response: UsersResponse;
	create: UsersCreate;
	update: UsersUpdate;
	relations: {
		'teams(users)': TeamsCollection[];
	};
};

// ===== teams =====

export type TeamsResponse = {
    users?: Array<string>;
} & BaseCollectionRecord;

export type TeamsCreate = {
	users?: MaybeArray<string>;
};

export type TeamsUpdate = {
	users?: MaybeArray<string>;
	'users+'?: MaybeArray<string>;
	'users-'?: MaybeArray<string>;
};

export type TeamsCollection = {
	type: 'base';
	collectionId: 'qtgju6qow8zu5sc';
	collectionName: 'teams';
	response: TeamsResponse;
	create: TeamsCreate;
	update: TeamsUpdate;
	relations: {
		users: UsersCollection[];
	};
};

// ===== projects =====

export type ProjectsResponse = {
    repoURL?: string;
	name?: string;
	deployed?: boolean;
	buildpack?: 'nixpacks' | 'dockerfile' | 'docker-compose';
	configured?: boolean;
	baseDir?: string;
	buildDir?: string;
	https?: boolean;
	www?: boolean;
	application?: any;
	managed?: boolean;
} & BaseCollectionRecord;

export type ProjectsCreate = {
	repoURL?: string | URL;
	name?: string;
	deployed?: boolean;
	buildpack?: 'nixpacks' | 'dockerfile' | 'docker-compose';
	configured?: boolean;
	baseDir?: string;
	buildDir?: string;
	https?: boolean;
	www?: boolean;
	application?: any;
	managed?: boolean;
};

export type ProjectsUpdate = {
	repoURL?: string | URL;
	name?: string;
	deployed?: boolean;
	buildpack?: 'nixpacks' | 'dockerfile' | 'docker-compose';
	configured?: boolean;
	baseDir?: string;
	buildDir?: string;
	https?: boolean;
	www?: boolean;
	application?: any;
	managed?: boolean;
};

export type ProjectsCollection = {
	type: 'base';
	collectionId: 'tqlcamhhas2xzr7';
	collectionName: 'projects';
	response: ProjectsResponse;
	create: ProjectsCreate;
	update: ProjectsUpdate;
	relations: {
		'logs(project)': LogsCollection[];
        'queue(project)': QueueCollection[];
	};
};

// ===== logs =====

export type LogsResponse = {
    project?: Array<string>;
} & BaseCollectionRecord;

export type LogsCreate = {
	project?: MaybeArray<string>;
};

export type LogsUpdate = {
	project?: MaybeArray<string>;
	'project+'?: MaybeArray<string>;
	'project-'?: MaybeArray<string>;
};

export type LogsCollection = {
	type: 'base';
	collectionId: '6po51li0eihk83d';
	collectionName: 'logs';
	response: LogsResponse;
	create: LogsCreate;
	update: LogsUpdate;
	relations: {
		project: ProjectsCollection[];
	};
};

// ===== templates =====

export type TemplatesResponse = {
    type?: 'portainer' | 'caprover' | 'custom';
	content?: string;
} & BaseCollectionRecord;

export type TemplatesCreate = {
	type?: 'portainer' | 'caprover' | 'custom';
	content?: string;
};

export type TemplatesUpdate = {
	type?: 'portainer' | 'caprover' | 'custom';
	content?: string;
};

export type TemplatesCollection = {
	type: 'base';
	collectionId: '67nwlz2znzqld2l';
	collectionName: 'templates';
	response: TemplatesResponse;
	create: TemplatesCreate;
	update: TemplatesUpdate;
	relations: {};
};

// ===== git =====

export type GitResponse = {
    app_id?: string;
	webhook_secret?: string;
	private_key?: string;
} & BaseCollectionRecord;

export type GitCreate = {
	app_id?: string;
	webhook_secret?: string;
	private_key?: string;
};

export type GitUpdate = {
	app_id?: string;
	webhook_secret?: string;
	private_key?: string;
};

export type GitCollection = {
	type: 'base';
	collectionId: '5c1oykad2csm1dk';
	collectionName: 'git';
	response: GitResponse;
	create: GitCreate;
	update: GitUpdate;
	relations: {};
};

// ===== queue =====

export type QueueResponse = {
    project?: string;
	active?: boolean;
} & BaseCollectionRecord;

export type QueueCreate = {
	project?: string;
	active?: boolean;
};

export type QueueUpdate = {
	project?: string;
	active?: boolean;
};

export type QueueCollection = {
	type: 'base';
	collectionId: 'ojmqge2b4cj4ywj';
	collectionName: 'queue';
	response: QueueResponse;
	create: QueueCreate;
	update: QueueUpdate;
	relations: {
		project: ProjectsCollection;
	};
};

// ===== Schema =====

export type Schema = {
	users: UsersCollection;
	teams: TeamsCollection;
	projects: ProjectsCollection;
	logs: LogsCollection;
	templates: TemplatesCollection;
	git: GitCollection;
	queue: QueueCollection;
};