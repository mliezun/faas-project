CREATE TABLE users (
    userId CHAR(32) PRIMARY KEY NOT NULL DEFAULT UUID(),
    userAttributes JSON NOT NULL
) ENGINE=RocksDB;

CREATE TABLE powers (
    userId CHAR(32) NOT NULL,
    projectId CHAR(32) NOT NULL,
    powersGranted JSON NOT NULL, -- default: ALL
    PRIMARY KEY (userId, projectId)
) ENGINE=RocksDB;

CREATE TABLE projects (
    projectId CHAR(32) PRIMARY KEY NOT NULL DEFAULT UUID(),
    projectName CHAR(100) NOT NULL,
    projectAttributes JSON NOT NULL,
    INDEX(projectName)
) ENGINE=RocksDB;

CREATE TABLE domains (
    domainId CHAR(32) PRIMARY KEY NOT NULL DEFAULT UUID(),
    projectId CHAR(32) NOT NULL, -- relation with project
    domain VARCHAR(1000) NOT NULL,
    domainStatus INT NOT NULL,
    domainAttributes JSON NOT NULL
    INDEX(domain)
) ENGINE=RocksDB;

CREATE TABLE collections (
    collectionId CHAR(32) PRIMARY KEY NOT NULL DEFAULT UUID(),
    projectId CHAR(32) NOT NULL,
    collectionName VARCHAR(256) NOT NULL,
    collectionAttributes JSON NOT NULL,
    INDEX(projectId, collectionName)
) ENGINE=RocksDB;

CREATE TABLE records (
    sequenceId BIGINT NOT NULL PRIMARY KEY, -- vitess_sequence
    recordId CHAR(32) NOT NULL, -- relation with project
    recordKey VARBINARY(1024) NOT NULL,
    recordValue JSON NOT NULL, -- any json value. ex: "literal string"
    created_at TIMESTAMP(6) NOT NULL,
    expiration TIMESTAMP(6),
    recordAttributes JSON NOT NULL,
    INDEX(keyStore, projectId, sequenceId DESC), -- find last valid element by key
    INDEX(projectId, keyStore, expiration) -- find all keys for a project
) ENGINE=RocksDB;

CREATE TABLE functions (
    functionId CHAR(32) PRIMARY KEY NOT NULL DEFAULT UUID(),
    functionBody TEXT NOT NULL,
    compilerVersion CHAR(11) NOT NULL, -- max: '999.999.999'
    created_at TIMESTAMP(6) NOT NULL,
    compiledCode BLOB NOT NULL,
    functionAttributes json
) ENGINE=RocksDB;
