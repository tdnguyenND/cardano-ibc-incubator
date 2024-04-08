import {
  EVENT_TYPE_CONNECTION,
  ATTRIBUTE_KEY_CONNECTION,
  STATE_MAPPING_CONNECTION,
  CONNECTION_ID_PREFIX,
} from './connection';
import {
  EVENT_TYPE_CHANNEL,
  ATTRIBUTE_KEY_CHANNEL,
  CHANNEL_ID_PREFIX,
  PORT_ID_PREFIX,
  TRANSFER_MODULE_PORT,
} from './channel';
import { EVENT_TYPE_CLIENT, ATTRIBUTE_KEY_CLIENT, CLIENT_ID_PREFIX } from './client';
import { EVENT_TYPE_SPO } from './spo';
import { REDEEMER_EMPTY_DATA, REDEEMER_TYPE } from './redeemer';
import { EVENT_TYPE_PACKET, ATTRIBUTE_KEY_PACKET } from './packet';
import {
  BLOCKID_FLAG_ABSENT,
  BLOCKID_FLAG_COMMIT,
  BLOCKID_FLAG_NIL,
  CRYPTO_ADDRESS_SIZE,
  MAX_SIGNATURE_SIZE,
  MAX_CHAIN_ID_LENGTH,
  TM_HASH_SIZE,
} from './block';

export const HANDLER_TOKEN_NAME = '68616e646c6572'; // fromText("handler")
export const CLIENT_PREFIX = '6962635f636c69656e74'; // fromText("ibc_client")
export const CONNECTION_TOKEN_PREFIX = '636f6e6e656374696f6e'; // fromText("connection")
export const CHANNEL_TOKEN_PREFIX = '6368616e6e656c'; // fromText("channel")
export const DEFAULT_IDENTIFIER_VERSION = '31'; // fromText("1")
export const DEFAULT_FEATURES_VERSION_ORDER_ORDERED = '4f524445525f4f524445524544'; // fromText("ORDER_ORDERED")
export const DEFAULT_FEATURES_VERSION_ORDER_UNORDERED = '4f524445525f554e4f524445524544'; // fromText("connection")
export const ACK_RESULT = '01';

export const LOVELACE = 'lovelace';

export const EMULATOR_ENV = 'emulator';
export {
  EVENT_TYPE_CONNECTION,
  ATTRIBUTE_KEY_CONNECTION,
  STATE_MAPPING_CONNECTION,
  CONNECTION_ID_PREFIX,
  EVENT_TYPE_CHANNEL,
  ATTRIBUTE_KEY_CHANNEL,
  CHANNEL_ID_PREFIX,
  PORT_ID_PREFIX,
  TRANSFER_MODULE_PORT,
  REDEEMER_TYPE,
  REDEEMER_EMPTY_DATA,
  EVENT_TYPE_CLIENT,
  ATTRIBUTE_KEY_CLIENT,
  CLIENT_ID_PREFIX,
  EVENT_TYPE_SPO,
  EVENT_TYPE_PACKET,
  ATTRIBUTE_KEY_PACKET,
  // block
  BLOCKID_FLAG_ABSENT,
  BLOCKID_FLAG_COMMIT,
  BLOCKID_FLAG_NIL,
  CRYPTO_ADDRESS_SIZE,
  MAX_SIGNATURE_SIZE,
  MAX_CHAIN_ID_LENGTH,
  TM_HASH_SIZE,
};
