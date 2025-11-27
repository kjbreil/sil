# SIL Library Specification

A comprehensive specification for the SIL (Standard Interchange Language) file format library, designed for retail data exchange in point-of-sale systems.

## Table of Contents

1. [Overview](#overview)
2. [SIL File Format](#sil-file-format)
3. [Core Types](#core-types)
4. [Tag System](#tag-system)
5. [Writer Module](#writer-module)
6. [Reader Module](#reader-module)
7. [Predefined Table Types](#predefined-table-types)
8. [Utilities](#utilities)
9. [Platform-Specific Features](#platform-specific-features)
10. [Example Usage](#example-usage)

---

## Overview

### Purpose

The SIL library provides bidirectional marshaling capabilities for SIL (Standard Interchange Language) files, a text-based format used in retail systems for data exchange. It allows:

- **Writing**: Convert structured data objects into SIL file format
- **Reading**: Parse SIL files back into structured data objects

### Key Features

- Reflection-based serialization using struct tags
- Support for optional fields (pointers)
- Julian date formatting for retail systems
- Windows-1252 character encoding for compatibility
- Streaming parser with channel-based data delivery
- Multi-batch file support
- Field grouping optimization for compact output

---

## SIL File Format

### Structure

A SIL file consists of three main sections:

```
1. HEADER (optional, except for LOAD action)
2. CREATE VIEW + INSERT statements with DATA rows
3. FOOTER (optional custom strings)
```

### Example File

```sql
INSERT INTO HEADER_DCT VALUES
('HM','12345678','MANUAL','PAL','','',2024365,0000,2024365,0000,'','ADDRPL','ADDRPL FROM GO','','','','','','','','','');

CREATE VIEW OBJ_CHG AS SELECT F01,F1000,F1001,F253 FROM OBJ_DCT;

INSERT INTO OBJ_CHG VALUES
('0001234567890','PAL',1,2024365),
('0009876543210','PAL',1,2024365);
```

### Format Rules

| Rule | Description |
|------|-------------|
| Line Endings | CRLF (`\r\n`) |
| Encoding | Windows-1252 (CP1252) |
| String Values | Enclosed in single quotes: `'value'` |
| Escaped Quotes | Single quotes escaped as doubled: `''` |
| Integer Values | Unquoted numbers |
| Statement End | Semicolon (`;`) |
| Data Rows | Comma-separated within parentheses: `(v1,v2,v3)` |
| Row Separator | Comma between rows, none before final semicolon |

### Table Naming Convention

Tables follow the pattern: `TABLENAME_ACTION`

| Action | Description |
|--------|-------------|
| `DCT` | Dictionary/Definition |
| `CHG` | Change (default for updates) |
| `LOAD` | Full load operation |
| `RSP` | Response |

### Field Codes

Fields are identified by F-codes (e.g., `F01`, `F902`, `F1000`). These are the SIL standard field identifiers used across retail systems.

---

## Core Types

### SIL

The main structure representing a SIL file.

```typescript
interface SIL {
  header: Header;
  view: View;
  footer: Footer;
  prefix: number;        // Internal: random batch prefix (0-99)
  tableType: unknown;    // Type definition for reflection
  include: boolean;      // Include empty optional fields
}
```

### Header

Batch metadata following the SIL header specification.

```typescript
interface Header {
  type: string;              // F901 - Batch type (default: "HM")
  identifier: string;        // F902 - Batch identifier (default: "00000001")
  creator: string;           // F903 - Batch creator (default: "MANUAL")
  destination: string;       // F904 - Batch destination (default: "PAL")
  auditFile: string;         // F905 - Batch audit file
  responseFile: string;      // F906 - Batch response file
  endingDate: number;        // F907 - Batch ending date (default: NOW)
  endingTime: number;        // F908 - Batch ending time (default: 0000)
  activeDate: number;        // F909 - Batch active date (default: NOW)
  activeTime: number;        // F910 - Batch active time (default: 0000)
  purgeDate: string;         // F911 - Batch purge date
  actionType: string;        // F912 - Batch action type (default: "ADDRPL")
  description: string;       // F913 - Batch description (default: "ADDRPL FROM GO")
  userOneState: string;      // F914 - Batch user 1 (state)
  maximumErrorCount: string; // F918 - Batch maximum error count
  fileVersion: string;       // F919 - Batch file version
  creatorVersion: string;    // F920 - Batch creator version
  primaryKey: string;        // F921 - Batch primary key
  specificCommand: string;   // F922 - Batch specific command
  tagType: string;           // F930 - Shelf tag type
  executionPriority: string; // F931 - Batch execution priority
  longDescription: string;   // F932 - Batch long description
}
```

### Header Action Types

```typescript
type ActionType = 'ADD' | 'ADDRPL' | 'CHANGE' | 'REMOVE';

const ActionTypes = {
  ADD: 'ADD',
  ADDRPL: 'ADDRPL',
  CHANGE: 'CHANGE',
  REMOVE: 'REMOVE'
} as const;
```

### View

Contains the table data.

```typescript
interface View {
  name: string;          // Table name (e.g., "OBJ", "PRICE")
  required: boolean;     // Whether view is required
  action: string;        // Action suffix (default: "CHG")
  data: unknown[];       // Array of data rows (any struct type)
}
```

### Footer

Optional content appended to the end of the file.

```typescript
type Footer = string[];
```

### Multi

A collection of multiple SIL batches (for multi-table files).

```typescript
type Multi = Map<string, SIL>;
```

---

## Tag System

### Struct Tag Format

Data structures use tags to define SIL field mappings:

```typescript
// TypeScript decorator equivalent concept
interface FieldMetadata {
  sil: string;      // Required: Field code (e.g., "F01")
  options?: string[]; // Optional: "zeropad"
  default?: string;  // Optional: default value or "NOW"/"NOWTIME"
  json?: string;     // Optional: JSON export name
}
```

### Tag Options

| Tag | Description | Example |
|-----|-------------|---------|
| `sil` | Field code (required) | `sil:"F01"` |
| `sil` + `zeropad` | Zero-pad to 13 characters | `sil:"F01,zeropad"` |
| `default` | Default value | `default:"PAL"` |
| `default:"NOW"` | Current Julian date | `default:"NOW"` |
| `default:"NOWTIME"` | Current Julian datetime | `default:"NOWTIME"` |

### Field Requirements

| Go Type | Behavior |
|---------|----------|
| Non-pointer | Required field - error if empty |
| Pointer | Optional field - omitted if nil (unless `include=true`) |

### SilTag Internal Structure

```typescript
interface SilTag {
  name: string;                    // Field code (e.g., "F01")
  options: string[];               // Additional options (e.g., ["zeropad"])
  field: ReflectionStructField;    // Reference to the struct field
}
```

### Tag Parsing Logic

```typescript
function getSilTag(field: StructField): { tag: SilTag; pad: boolean } {
  const silTagStr = field.tags.get('sil');
  if (!silTagStr) {
    throw new Error('does not contain a sil tag');
  }

  const parts = silTagStr.split(',');
  const tag: SilTag = {
    name: parts[0],
    options: [],
    field: field
  };

  // Check for zeropad option
  let pad = false;
  for (let i = 1; i < parts.length; i++) {
    if (parts[i] === 'zeropad') {
      tag.options.push(parts[i]);
      pad = true;
    }
  }

  return { tag, pad };
}
```

---

## Writer Module

### Core Functions

#### Make

Creates a new SIL structure for a specific table.

```typescript
function make(name: string, definition: unknown): SIL {
  return {
    header: getDefaultHeader(),
    view: {
      name: name,
      required: true,
      action: '',
      data: []
    },
    footer: [],
    prefix: Math.floor(Math.random() * 100),
    tableType: definition,
    include: false
  };
}
```

#### Marshal

Converts the SIL structure to bytes.

```typescript
async function marshal(sil: SIL, include: boolean): Promise<Uint8Array> {
  // 1. Validate view name is set
  if (!sil.view.name) {
    throw new Error('view name not set');
  }

  // 2. Override include with SIL's include if true
  if (!include) {
    include = sil.include;
  }

  // 3. Split data into sections by field signature
  const sections = split(sil.view.data, include);

  // 4. Build output
  let data: string[] = [];

  for (const section of sections.values()) {
    // Generate batch identifier if not set
    if (!sil.header.identifier) {
      sil.header.identifier = batchNum(sil.prefix);
    }

    // Add header (unless LOAD action)
    if (sil.view.action !== 'LOAD') {
      data.push(headerInsert());
      data.push(headerRow(sil.header));
    }

    // Add section data
    data.push(sectionCreate(section, sil.view));
    data.push('\r\n');
  }

  // 5. Add footer
  data.push(sil.footer.join(''));

  // 6. Encode to Windows-1252
  return encodeWindows1252(data.join(''));
}
```

### Section Processing

#### Split Function

Groups rows by their field signature (combination of present fields).

```typescript
function split(rows: unknown[], include: boolean): Map<string, Section> {
  const sections = new Map<string, Section>();

  for (const row of rows) {
    const rowData = makeRow(row, include);

    // Create key from field names
    const key = rowData.elems.map(e => e.name).join('');

    if (!sections.has(key)) {
      sections.set(key, []);
    }
    sections.get(key)!.push(rowData);
  }

  return sections;
}
```

#### Row Structure

```typescript
interface Elem {
  name: string | null;   // Field code (e.g., "F01")
  data: string | null;   // Formatted value
}

interface Row {
  elems: Elem[];
}

type Section = Row[];
```

#### Section Create

Generates the CREATE VIEW and INSERT statements.

```typescript
function sectionCreate(section: Section, view: View): string {
  // Get field names from first row
  const names = section[0].elems.map(e => e.name);
  const namesStr = names.join(',');

  const actionName = view.action ? `${view.name}_${view.action}` : `${view.name}_CHG`;

  let output = `CREATE VIEW ${actionName} AS SELECT ${namesStr} FROM ${view.name}_DCT;\r\n\r\n`;
  output += `INSERT INTO ${actionName} VALUES\r\n`;

  // Add data rows
  const rows = section.map(row => {
    const values = row.elems.map(e => e.data);
    return `(${values.join(',')})`;
  });

  output += rows.join('\r\n');
  output += ';\r\n';

  return output;
}
```

### Value Formatting

#### Value Function

Extracts and formats a field value.

```typescript
function formatValue(
  value: unknown,
  field: FieldMetadata
): { value: string; name: string; isPointer: boolean } {
  const tag = getSilTag(field);
  const defaultValue = getDefaultTag(field);

  const { formatted, isPointer } = formatByKind(value, defaultValue);

  // Apply zeropad if needed
  let result = formatted;
  if (tag.pad && result.length > 0) {
    // Remove quotes, pad, re-add quotes
    const inner = result.slice(1, -1);
    result = `'${inner.padStart(13, '0')}'`;

    // Validate length (15 chars including quotes)
    if (result.length > 15) {
      throw new Error(`padded field contains more than 13 characters: ${result}`);
    }
  }

  return {
    value: result,
    name: tag.name,
    isPointer
  };
}
```

#### Kind Formatting

```typescript
function formatByKind(
  value: unknown,
  defaultValue: string
): { formatted: string; isPointer: boolean } {
  const hasDefault = defaultValue.length > 0;

  // Handle null/undefined (pointer case)
  if (value === null || value === undefined) {
    return { formatted: '', isPointer: true };
  }

  // Handle string
  if (typeof value === 'string') {
    if (value.length === 0 && hasDefault) {
      return { formatted: `'${defaultValue}'`, isPointer: false };
    }
    if (value.length === 0) {
      return { formatted: '', isPointer: false };
    }
    // Escape single quotes
    const escaped = value.replace(/'/g, "''");
    return { formatted: `'${escaped}'`, isPointer: false };
  }

  // Handle number (integer)
  if (typeof value === 'number') {
    if (value === 0 && hasDefault) {
      return { formatted: defaultValue, isPointer: false };
    }
    return { formatted: String(Math.floor(value)), isPointer: false };
  }

  return { formatted: '', isPointer: false };
}
```

#### Default Tag Processing

```typescript
function getDefaultTag(field: FieldMetadata): string {
  const def = field.default;

  if (!def) return '';

  switch (def) {
    case 'NOW':
      return julianNow();
    case 'NOWTIME':
      return julianTimeNow();
    default:
      return def;
  }
}
```

### Multi-Batch Support

```typescript
class Multi {
  private batches: Map<string, SIL> = new Map();

  make(name: string, definition: unknown): void {
    this.batches.set(name, make(name, definition));
  }

  appendView(name: string, data: unknown): void {
    const sil = this.batches.get(name);
    if (sil) {
      sil.view.data.push(data);
    }
  }

  setHeaders(description: string): void {
    for (const sil of this.batches.values()) {
      sil.header.description = description;
    }
  }

  async marshal(): Promise<Uint8Array> {
    const prefix = Math.floor(Math.random() * 100);
    const chunks: Uint8Array[] = [];

    for (const sil of this.batches.values()) {
      sil.prefix = prefix;
      const bytes = await marshal(sil, false);
      chunks.push(bytes);
    }

    return concatenateArrays(chunks);
  }
}
```

### File Writing

```typescript
async function write(
  sil: SIL,
  filename: string,
  include: boolean,
  archive: boolean
): Promise<void> {
  const data = await marshal(sil, include);

  // Set archive bit before writing (Windows only)
  if (archive) {
    await setArchive(filename);
  }

  await writeFile(filename, data);

  // Unset archive bit after writing (Windows only)
  if (archive) {
    await unsetArchive(filename);
  }
}
```

### Batch Number Generation

```typescript
function batchNum(prefix: number): string {
  const random = Math.floor(Math.random() * 1000000);
  return `${prefix.toString().padStart(2, '0')}${random.toString().padStart(6, '0')}`;
}
```

---

## Reader Module

### Token Types

```typescript
enum Token {
  ILLEGAL = 0,
  EOF = 1,
  WS = 2,        // Whitespace (space, tab)
  IDENT = 3,     // Identifiers/field names
  COMMA = 4,     // ,
  OPEN = 5,      // (
  CLOSE = 6,     // )
  SEMICOLON = 7, // ;
  SINGLE = 8,    // '
  CRLF = 9       // \r\n
}
```

### Part (Token + Value)

```typescript
interface Part {
  tok: Token;
  val: string;
}

type Parsed = Part[];
```

### Scanner

Lexical analysis - converts character stream to tokens.

```typescript
class Scanner {
  private reader: BufferedReader;

  constructor(input: Reader) {
    this.reader = new BufferedReader(input);
  }

  scan(): Part {
    const ch = this.read();

    // Check for whitespace
    if (this.isWhitespace(ch)) {
      this.unread();
      return this.scanWhitespace();
    }

    // Check for specific characters
    switch (ch) {
      case '\0': // EOF
        return { tok: Token.EOF, val: '' };
      case ',':
        return { tok: Token.COMMA, val: ',' };
      case '(':
        return { tok: Token.OPEN, val: '(' };
      case ')':
        return { tok: Token.CLOSE, val: ')' };
      case ';':
        return { tok: Token.SEMICOLON, val: ';' };
      case "'":
        return { tok: Token.SINGLE, val: "'" };
      case '\r':
        const next = this.read();
        if (next !== '\n') {
          throw new Error('carriage return without newline');
        }
        return { tok: Token.CRLF, val: '\r\n' };
      case '\n':
        // Newline without carriage return - treat as CRLF
        return { tok: Token.CRLF, val: '\r\n' };
    }

    this.unread();
    return this.scanIdent();
  }

  private scanIdent(): Part {
    let buf = this.read();

    while (true) {
      const ch = this.read();
      if (ch === '\0') break;
      if (!this.isLetter(ch) && !this.isDigit(ch) && !this.isIncludeSpecial(ch)) {
        this.unread();
        break;
      }
      buf += ch;
    }

    return { tok: Token.IDENT, val: buf };
  }

  private scanWhitespace(): Part {
    let buf = this.read();

    while (true) {
      const ch = this.read();
      if (ch === '\0' || !this.isWhitespace(ch)) {
        this.unread();
        break;
      }
      buf += ch;
    }

    return { tok: Token.WS, val: buf };
  }

  private isWhitespace(ch: string): boolean {
    return ch === ' ' || ch === '\t';
  }

  private isLetter(ch: string): boolean {
    return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z');
  }

  private isDigit(ch: string): boolean {
    return ch >= '0' && ch <= '9';
  }

  private isIncludeSpecial(ch: string): boolean {
    return ['_', '.', '+', '/', ':'].includes(ch);
  }
}
```

### Parser

Wraps scanner with lookahead buffer.

```typescript
class Parser {
  private scanner: Scanner;
  private buffer: { pt: Part; n: number } = { pt: { tok: Token.EOF, val: '' }, n: 0 };

  constructor(input: Reader) {
    this.scanner = new Scanner(input);
  }

  scan(): Part {
    // Return buffered token if available
    if (this.buffer.n !== 0) {
      this.buffer.n = 0;
      return this.buffer.pt;
    }

    // Otherwise scan next token
    const pt = this.scanner.scan();
    this.buffer.pt = pt;
    return pt;
  }

  parse(): Parsed {
    const parsed: Parsed = [];

    while (true) {
      const pt = this.scan();
      parsed.push(pt);
      if (pt.tok === Token.EOF) break;
    }

    return parsed;
  }

  advanceTo(tok: Token): void {
    while (this.scan().tok !== tok) {
      // Keep scanning
    }
  }
}
```

### Decoder

Semantic analysis - interprets parsed tokens.

```typescript
interface Decoder {
  p: Parsed;
  err: Error[];
  fcodes: string[];
  fieldMap: number[];
  tableName: string;
  view: boolean;         // Has reached view data
  headerInsert: boolean; // Next line is header info
  header: string[];
  data: string[][];
  macroStrings: StringBuilder;
}

function newDecoder(): Decoder {
  return {
    p: [],
    err: [],
    fcodes: [],
    fieldMap: [],
    tableName: '',
    view: false,
    headerInsert: false,
    header: [],
    data: [],
    macroStrings: new StringBuilder()
  };
}
```

#### Line Identification

```typescript
function identifyLine(d: Decoder, s: number): number {
  // If headerInsert flag is set, read header data
  if (d.headerInsert) {
    const e = nextCRLF(d.p, s);

    // Skip HC (header continuation) lines
    if (d.p[s + 2]?.val === 'HC') {
      return nextCRLF(d.p, s);
    }

    // Validate header row format
    if (d.p[s].tok !== Token.OPEN ||
        d.p[e - 2].tok !== Token.CLOSE ||
        d.p[e - 1].tok !== Token.SEMICOLON) {
      d.err.push(new Error('row for HEADER invalid'));
      return s;
    }

    const result = readDataLine(d.p, s, 22);
    d.header = result.data;
    d.headerInsert = false;
    return result.nextIndex;
  }

  // If view has been reached, read data rows
  if (d.view) {
    const result = readDataLine(d.p, s, d.fcodes.length);
    if (result.error) {
      d.err.push(result.error);
    }
    d.data.push(result.data);
    return result.nextIndex;
  }

  // Identify line type by first token
  switch (d.p[s].tok) {
    case Token.CRLF:
      return s + 1;
    case Token.OPEN:
      return readInsertLine(d, s);
  }

  // Identify by value
  switch (d.p[s].val) {
    case 'INSERT':
      return checkInsert(d, s);
    case 'CREATE':
      return checkCreate(d, s);
  }

  // Skip unrecognized line
  while (d.p[s].tok !== Token.CRLF) {
    s++;
  }
  return s;
}
```

#### Data Line Reading

```typescript
interface ReadResult {
  data: string[];
  nextIndex: number;
  error?: Error;
}

function readDataLine(p: Parsed, s: number, columns: number): ReadResult {
  const lineData: string[] = [];

  // Expect opening parenthesis
  if (p[s].tok !== Token.OPEN) {
    return { data: lineData, nextIndex: s, error: new Error('data does not start with (') };
  }
  s++;

  // Read each column
  for (let i = 0; i < columns; i++) {
    const { data, nextIndex } = readData(p, s);
    s = nextIndex;

    // Check for comma (except last column)
    if (p[s].tok !== Token.COMMA && i !== columns - 1 && data !== '') {
      return { data: lineData, nextIndex: s, error: new Error('data does not end with ,') };
    } else if (p[s].tok === Token.COMMA && data !== '') {
      s++;
    }

    lineData.push(data);
  }

  // Expect closing parenthesis
  if (p[s].tok === Token.CLOSE) {
    s++;
  } else {
    return { data: lineData, nextIndex: s, error: new Error('data does not end with )') };
  }

  // Handle end of data
  if (p[s].tok === Token.SEMICOLON) {
    s += 2; // Skip semicolon and CRLF
    return { data: lineData, nextIndex: s };
  }

  // Handle comma between rows
  if (p[s].tok === Token.COMMA) {
    s++;
  }

  // Handle end of line
  if (p[s].tok === Token.CRLF) {
    s++;
  } else {
    return { data: lineData, nextIndex: s, error: new Error('no endline at end of data') };
  }

  return { data: lineData, nextIndex: s };
}
```

#### Single Value Reading (with quote handling)

```typescript
function readData(p: Parsed, s: number): { data: string; nextIndex: number } {
  let single = false; // Inside single quotes
  let data = '';
  let opens = 0;      // Nested parenthesis count

  while (true) {
    if (single) {
      if (p[s].tok === Token.SINGLE) {
        single = false;
        // Check for escaped quote (two single quotes)
        if (p[s + 1]?.tok === Token.SINGLE) {
          s++;
          single = true;
          data += "''";
        }
      } else {
        data += p[s].val;
      }
    } else {
      switch (p[s].tok) {
        case Token.SINGLE:
          single = true;
          // Check for escaped quote at start
          if (p[s + 1]?.tok === Token.SINGLE) {
            s++;
            single = false;
            data += "''";
          }
          break;
        case Token.OPEN:
          opens++;
          break;
        case Token.CLOSE:
          opens--;
          if (opens < 0) {
            return { data, nextIndex: s };
          }
          break;
        case Token.COMMA:
          if (opens === 0) {
            // Empty field - advance past comma
            if (data.length === 0) {
              return { data, nextIndex: s + 1 };
            }
            return { data, nextIndex: s };
          }
          data += p[s].val;
          break;
        default:
          data += p[s].val;
      }
    }
    s++;
  }
}
```

### Unmarshal Functions

#### Unmarshal Bytes

```typescript
async function unmarshal<T>(bytes: Uint8Array, dataType: new () => T): Promise<T[]> {
  const reader = new ByteReader(bytes);

  // Create channel for async data delivery
  const channel = new Channel<T>(100);

  await unmarshalReaderChan(reader, channel);

  // Collect all results
  const results: T[] = [];
  for await (const item of channel) {
    results.push(item);
  }

  return results;
}
```

#### Unmarshal to Channel

```typescript
async function unmarshalReaderChan<T>(
  reader: Reader,
  channel: Channel<T>
): Promise<void> {
  const parser = new Parser(reader);

  // Process in background
  processDecoder(parser, channel);
}

async function processDecoder<T>(
  parser: Parser,
  channel: Channel<T>
): Promise<void> {
  const decoder = newDecoder();
  let i = 0;

  try {
    while (true) {
      const pt = parser.scan();
      decoder.p.push(pt);

      if (pt.tok === Token.CRLF) {
        const newIndex = identifyLine(decoder, i);

        // If index didn't advance, something went wrong
        if (newIndex === i) break;

        i = 0;
        decoder.p = [];

        // If view data found, process and send
        if (decoder.view) {
          for (const data of decoder.data) {
            if (decoder.fieldMap.length === 0) {
              makeFieldMap(decoder, channel.type);
            }

            const result = unmarshalValues<T>(data, decoder.fieldMap);
            if (result) {
              channel.send(result);
            }
          }
          decoder.data = [];
        }
      }

      if (pt.tok === Token.EOF) break;
    }
  } finally {
    channel.close();
  }
}
```

#### Value Unmarshaling

```typescript
function unmarshalValues<T>(
  input: string[],
  fieldMap: number[],
  resultType: new () => T
): T {
  if (fieldMap.length === 0) {
    throw new Error('fieldMap is empty');
  }

  const result = new resultType();

  for (let c = 0; c < input.length; c++) {
    const fieldIndex = fieldMap[c];
    if (fieldIndex === -1) continue;
    if (input[c] === '') continue;

    unmarshalValue(result, fieldIndex, input[c]);
  }

  return result;
}

function unmarshalValue(
  target: object,
  fieldIndex: number,
  input: string
): void {
  const field = getFieldByIndex(target, fieldIndex);
  const fieldType = getFieldType(target, fieldIndex);

  switch (fieldType) {
    case 'string':
      setField(target, fieldIndex, input);
      break;

    case 'number':
      const num = input.length === 0 ? 0 : parseInt(input, 10);
      if (isNaN(num)) {
        throw new Error(`conversion of int failed from: ${input}`);
      }
      setField(target, fieldIndex, num);
      break;

    case 'Date':
      // Julian date format: YYYYDDD or YYYYDDD HH:MM:SS
      let date: Date;
      if (input.length === 7) {
        date = parseJulianDate(input);
      } else if (input.length >= 16) {
        date = parseJulianDateTime(input.substring(0, 16));
      } else {
        return;
      }
      setField(target, fieldIndex, date);
      break;

    default:
      throw new Error(`unhandled type: ${fieldType}`);
  }
}
```

### Reader Class

```typescript
interface Stats {
  hasHeader: boolean;
  header: string[];
  hasView: boolean;
  table: string;
  dataLines: number;
  fcodes: string[];
  hasCreate: boolean;
}

class Reader {
  private reader: ReadSeekCloser;
  private parser: Parser;
  public stats: Stats;

  static async create(rs: ReadSeekCloser): Promise<Reader> {
    const reader = new Reader();
    reader.reader = rs;
    reader.parser = new Parser(rs);

    await reader.loadStats();

    return reader;
  }

  private async loadStats(): Promise<void> {
    this.stats = {
      hasHeader: false,
      header: [],
      hasView: false,
      table: '',
      dataLines: 0,
      fcodes: [],
      hasCreate: false
    };

    // Scan entire file to gather statistics
    // ... implementation similar to newStatsFromReader
  }

  async unmarshalChan<T>(channel: Channel<T>): Promise<void> {
    // Seek to beginning
    await this.reader.seek(0, SeekStart);

    const parser = new Parser(this.reader);
    processDecoder(parser, channel);
  }

  clearReader(): void {
    this.reader = null;
  }

  async close(): Promise<void> {
    await this.reader.close();
  }
}
```

### Field Mapping

```typescript
function makeFieldMap(decoder: Decoder, dataType: unknown): void {
  for (const fcode of decoder.fcodes) {
    const index = findFieldIndex(fcode, dataType);
    decoder.fieldMap.push(index);
  }
}

function findFieldIndex(fcode: string, type: unknown): number {
  const fields = getTypeFields(type);

  for (let i = 0; i < fields.length; i++) {
    const tag = getSilTag(fields[i]);
    if (tag === fcode) {
      return i;
    }
  }

  return -1; // Field not found
}
```

### Table Parsing Utilities

```typescript
function parseTable(text: string): { name: string; action: string } {
  const parts = text.toUpperCase().split('_');
  const action = parts[parts.length - 1];
  const name = parts.slice(0, -1).join('_');
  return { name, action };
}

function getTable(p: Parsed, s: number): string {
  const { name, action } = parseTable(p[s + 4].val);

  switch (action) {
    case 'DCT':
    case 'CHG':
    case 'RSP':
    case 'LOAD':
      return name;
    default:
      return 'ERROR';
  }
}

function getAction(p: Parsed, s: number): string {
  const { action } = parseTable(p[s + 4].val);
  return action;
}

function isInsert(p: Parsed, s: number, table: string): boolean {
  if (p[s].val !== 'INSERT') return false;
  if (p[s + 2].val !== 'INTO') return false;
  if (getTable(p, s).toUpperCase() !== table.toUpperCase()) return false;
  if (p[s + 6].val !== 'VALUES') return false;
  return true;
}

function nextCRLF(p: Parsed, s: number): number {
  for (let i = s; i < p.length; i++) {
    if (p[i].tok === Token.CRLF) {
      return i;
    }
  }
  return s;
}

function nextLine(p: Parsed, s: number): number {
  return nextCRLF(p, s) + 1;
}
```

---

## Predefined Table Types

### OBJ (Items)

```typescript
interface OBJ {
  upcCode: string;              // F01, zeropad - UPC code
  upcCodeFormat?: number;       // F07 - UPC format
  targetIdentifier: string;     // F1000, default: "PAL"
  recordStatus: number;         // F1001, default: 1
  lastChangeDate: string;       // F253, default: NOW
  expandedDescription?: string; // F29
  longDescription?: string;     // F255
  createdByUser?: number;       // F940
  modifiedByUser?: number;      // F941
}
```

### PRICE (Pricing)

```typescript
interface PRICE {
  upcCode: string;              // F01, zeropad
  targetIdentifier: string;     // F1000, default: "PAL"
  price?: string;               // F30
  priceQty?: string;            // F31
  priceMixmatch?: number;       // F32
  priceMethod?: string;         // F33
  batchIdentifier?: string;     // F902
  batchCreator?: string;        // F903
}
```

### DSS (File Installation)

```typescript
interface DSS {
  priority: number;       // F2727, default: 30 - Install priority
  author: string;         // F2728, default: "KYGL"
  option: string;         // F2729 - Option name
  destination: string;    // F2730 - Target folder
  script: string;         // F2731 - Filename
  fileDate: string;       // F2732
  lastChangeDate: string; // F253, default: "@DJSF @FMT(T6F,@NOW)"
  signature: string;      // F2733 - File hash
}
```

### CLL (Customers)

```typescript
interface CLL {
  targetIdentifier: string;          // F1000, default: "PAL"
  recordStatus: number;              // F1001, default: 1
  customerId: string;                // F1148
  alternateCustNumber?: string;      // F1577
  alternateCustType?: string;        // F1578
  maintenanceOperatorLevel?: number; // F1759
  mainAltCode?: string;              // F1898
  storeResponsible?: string;         // F1964
  lastChangeDate: string;            // F253, default: NOW
  batchIdentifier?: string;          // F902
}
```

### ECL (Item Links)

```typescript
interface ECL {
  upcCode: string;         // F01
  posDescription?: string; // F02
  itemLinkCode: string;    // F164, zeropad
}
```

---

## Utilities

### Julian Date Functions

Julian dates use the format `YYYYDDD` where DDD is the day of year (001-366).

```typescript
// Format: YYYYDDD (e.g., 2024365)
function julianDate(date: Date): string {
  const year = date.getFullYear();
  const dayOfYear = getDayOfYear(date);
  return `${year}${dayOfYear.toString().padStart(3, '0')}`;
}

// Format: YYYYDDD HH:MM:SS (e.g., 2024365 14:30:00)
function julianDateTime(date: Date): string {
  return `${julianDate(date)} ${julianTimePart(date)}`;
}

function julianNow(): string {
  return julianDate(new Date());
}

function julianTimeNow(): string {
  return julianDateTime(new Date());
}

function julianTimePart(date: Date): string {
  const h = date.getHours().toString().padStart(2, '0');
  const m = date.getMinutes().toString().padStart(2, '0');
  const s = date.getSeconds().toString().padStart(2, '0');
  return `${h}:${m}:${s}`;
}

// Helper
function getDayOfYear(date: Date): number {
  const start = new Date(date.getFullYear(), 0, 0);
  const diff = date.getTime() - start.getTime();
  const oneDay = 1000 * 60 * 60 * 24;
  return Math.floor(diff / oneDay);
}

// Parsing
function parseJulianDate(julian: string): Date {
  // Format: YYYYDDD
  const year = parseInt(julian.substring(0, 4), 10);
  const dayOfYear = parseInt(julian.substring(4, 7), 10);

  const date = new Date(year, 0, 1);
  date.setDate(dayOfYear);
  return date;
}

function parseJulianDateTime(julian: string): Date {
  // Format: YYYYDDD HH:MM:SS
  const date = parseJulianDate(julian.substring(0, 7));
  const timeParts = julian.substring(8).split(':');
  date.setHours(parseInt(timeParts[0], 10));
  date.setMinutes(parseInt(timeParts[1], 10));
  date.setSeconds(parseInt(timeParts[2], 10));
  return date;
}
```

### Windows-1252 Encoding

```typescript
// Encoding map for characters 128-255
const windows1252Map: { [key: number]: number } = {
  0x20AC: 0x80, // €
  0x201A: 0x82, // ‚
  0x0192: 0x83, // ƒ
  0x201E: 0x84, // „
  0x2026: 0x85, // …
  0x2020: 0x86, // †
  0x2021: 0x87, // ‡
  0x02C6: 0x88, // ˆ
  0x2030: 0x89, // ‰
  0x0160: 0x8A, // Š
  0x2039: 0x8B, // ‹
  0x0152: 0x8C, // Œ
  0x017D: 0x8E, // Ž
  0x2018: 0x91, // '
  0x2019: 0x92, // '
  0x201C: 0x93, // "
  0x201D: 0x94, // "
  0x2022: 0x95, // •
  0x2013: 0x96, // –
  0x2014: 0x97, // —
  0x02DC: 0x98, // ˜
  0x2122: 0x99, // ™
  0x0161: 0x9A, // š
  0x203A: 0x9B, // ›
  0x0153: 0x9C, // œ
  0x017E: 0x9E, // ž
  0x0178: 0x9F, // Ÿ
};

function encodeWindows1252(text: string): Uint8Array {
  const bytes: number[] = [];

  for (const char of text) {
    const code = char.charCodeAt(0);

    if (code < 128) {
      bytes.push(code);
    } else if (code < 256) {
      bytes.push(code);
    } else if (windows1252Map[code] !== undefined) {
      bytes.push(windows1252Map[code]);
    } else {
      bytes.push(0x3F); // '?' for unmappable characters
    }
  }

  return new Uint8Array(bytes);
}

function decodeWindows1252(bytes: Uint8Array): string {
  // Reverse mapping for decoding
  const reverseMap: { [key: number]: number } = {};
  for (const [unicode, win] of Object.entries(windows1252Map)) {
    reverseMap[win as unknown as number] = parseInt(unicode);
  }

  let result = '';
  for (const byte of bytes) {
    if (byte < 128) {
      result += String.fromCharCode(byte);
    } else if (reverseMap[byte] !== undefined) {
      result += String.fromCharCode(reverseMap[byte]);
    } else {
      result += String.fromCharCode(byte);
    }
  }

  return result;
}
```

---

## Platform-Specific Features

### Archive Bit (Windows Only)

Windows systems can manipulate the file archive attribute:

```typescript
// Windows implementation
async function setArchive(filename: string): Promise<void> {
  // Use Windows API to set FILE_ATTRIBUTE_ARCHIVE
  // syscall.SetFileAttributes(filename, FILE_ATTRIBUTE_ARCHIVE)
}

async function unsetArchive(filename: string): Promise<void> {
  // Get current attributes and remove archive bit
  // syscall.SetFileAttributes(filename, attrs & ~FILE_ATTRIBUTE_ARCHIVE)
}

// macOS/Linux implementation (no-op)
async function setArchive(filename: string): Promise<void> {
  // No-op on non-Windows platforms
}

async function unsetArchive(filename: string): Promise<void> {
  // No-op on non-Windows platforms
}
```

---

## Example Usage

### Creating a SIL File

```typescript
import { SIL, OBJ, make, marshal, write } from 'sil';

// Define data structure
interface MyItem {
  upcCode: string;           // F01, zeropad
  targetIdentifier: string;  // F1000
  recordStatus: number;      // F1001
  lastChangeDate: string;    // F253
  description?: string;      // F29 (optional)
}

// Create SIL
const sil = make('OBJ', {} as MyItem);

// Add data rows
sil.view.data.push({
  upcCode: '1234567890123',
  targetIdentifier: 'PAL',
  recordStatus: 1,
  lastChangeDate: '', // Will use NOW default
  description: 'Test Item'
});

sil.view.data.push({
  upcCode: '9876543210987',
  targetIdentifier: 'PAL',
  recordStatus: 1,
  lastChangeDate: ''
  // description omitted - will be excluded from output
});

// Set custom header values
sil.header.description = 'My Batch Import';
sil.header.actionType = 'ADDRPL';

// Marshal to bytes
const bytes = await marshal(sil, false);

// Or write directly to file
await write(sil, 'output.sil', false, false);
```

### Reading a SIL File

```typescript
import { Reader, unmarshal } from 'sil/silread';

// Define expected structure
interface MyItem {
  upcCode: string;
  targetIdentifier: string;
  recordStatus: number;
  description?: string;
}

// Method 1: Unmarshal bytes directly
const bytes = await readFile('input.sil');
const items = await unmarshal<MyItem>(bytes, MyItem);

for (const item of items) {
  console.log(`UPC: ${item.upcCode}, Status: ${item.recordStatus}`);
}

// Method 2: Use Reader for stats and streaming
const file = await openFile('input.sil');
const reader = await Reader.create(file);

console.log(`Table: ${reader.stats.table}`);
console.log(`Data lines: ${reader.stats.dataLines}`);
console.log(`Fields: ${reader.stats.fcodes.join(', ')}`);

// Stream data via channel
const channel = new Channel<MyItem>(100);
await reader.unmarshalChan(channel);

for await (const item of channel) {
  console.log(`Processing: ${item.upcCode}`);
}

await reader.close();
```

### Multi-Batch File

```typescript
import { Multi, OBJ, PRICE } from 'sil';

const multi = new Multi();

// Create multiple tables
multi.make('OBJ', {} as OBJ);
multi.make('PRICE', {} as PRICE);

// Add data to each
multi.appendView('OBJ', {
  upcCode: '1234567890123',
  targetIdentifier: 'PAL',
  recordStatus: 1,
  lastChangeDate: ''
});

multi.appendView('PRICE', {
  upcCode: '1234567890123',
  targetIdentifier: 'PAL',
  price: '9.99'
});

// Set common headers
multi.setHeaders('Multi-table Import');

// Write to file
await multi.write('multi.sil', false);
```

---

## Error Handling

### Writer Errors

| Error | Cause |
|-------|-------|
| `view name not set` | `SIL.View.Name` is empty |
| `element X does not contain data and is required` | Non-pointer field is empty |
| `padded field contains more than 13 characters` | Zeropad field exceeds limit |
| `does not contain a sil tag` | Struct field missing `sil` tag |
| `conversion to 1252 failed` | Character encoding error |

### Reader Errors

| Error | Cause |
|-------|-------|
| `data needs to be a pointer to a slice` | Wrong unmarshal target type |
| `data is not a slice` | Target is not a slice |
| `data needs to be a channel` | Channel unmarshal without channel |
| `data does not start with (` | Malformed data row |
| `data does not end with )` | Malformed data row |
| `no endline at end of data` | Missing CRLF |
| `fieldMap is empty` | No CREATE statement parsed |
| `conversion of int failed` | Invalid number format |

---

## Constants Reference

```typescript
// Line endings
const CRLF = '\r\n';

// Action types
const ActionType = {
  ADD: 'ADD',
  ADDRPL: 'ADDRPL',
  CHANGE: 'CHANGE',
  REMOVE: 'REMOVE'
} as const;

// Table actions
const TableAction = {
  DCT: 'DCT',
  CHG: 'CHG',
  RSP: 'RSP',
  LOAD: 'LOAD'
} as const;

// Default header values
const DefaultHeader = {
  type: 'HM',
  identifier: '00000001',
  creator: 'MANUAL',
  destination: 'PAL',
  endingTime: 0,
  activeTime: 0,
  actionType: 'ADDRPL',
  description: 'ADDRPL FROM GO'
} as const;

// Header has 22 fields (F901-F932)
const HEADER_FIELD_COUNT = 22;
```

---

## Implementation Notes for TypeScript

### Reflection Alternative

Since TypeScript doesn't have runtime reflection like Go, consider:

1. **Decorators** for field metadata
2. **Schema definition objects** alongside interfaces
3. **Code generation** from schema files

Example with decorators:

```typescript
function SilField(code: string, options?: { zeropad?: boolean; default?: string }) {
  return function (target: any, propertyKey: string) {
    // Store metadata
    Reflect.defineMetadata('sil', { code, ...options }, target, propertyKey);
  };
}

class OBJ {
  @SilField('F01', { zeropad: true })
  upcCode: string;

  @SilField('F1000', { default: 'PAL' })
  targetIdentifier: string;

  @SilField('F253', { default: 'NOW' })
  lastChangeDate: string;
}
```

### Async Considerations

- File I/O should be async
- Use async iterators for streaming
- Consider Web Streams API for browser compatibility

### Encoding

- Use `TextEncoder`/`TextDecoder` with polyfills for Windows-1252
- Or use libraries like `iconv-lite` for Node.js
