是的，**Go语言的xorm框架**支持通过**反向工程（reverse engineering）**功能，根据现有数据库表自动生成对应的Go结构体（struct）代码。这大大提高了开发效率，避免了手动编写模型的繁琐工作。

### 如何实现
xorm提供了一个独立的命令行工具（以前叫xorm tool，现在是**xorm.io/reverse**），专门用于从数据库表生成代码。

1. **安装工具**：
   ```
   go install xorm.io/reverse@latest
   ```

   （注意：该工具依赖CGO，可能需要安装相关开发环境。）

2. **基本使用方式**：
   - **命令行直接生成**（旧版常见方式）：
     示例（MySQL）：
     ```
     xorm reverse mysql "root:password@tcp(127.0.0.1:3306)/your_database?charset=utf8" templates/goxorm ./models
     ```
     这会根据数据库所有表生成带`xorm` tag的结构体代码，输出到`./models`目录。

   - **推荐新版方式：使用配置文件**（更灵活，支持指定表、映射规则等）：
     创建一个YAML配置文件（如`reverse.yml`）：
     ```yaml
     kind: reverse
     name: mydb
     source:
       database: mysql  # 或 postgres, sqlite3 等
       conn_str: root:password@tcp(127.0.0.1:3306)/your_database?charset=utf8
     targets:
       - type: codes
         language: golang
         output_dir: ./models
         multiple_files: true  # 每个表生成一个文件
         table_mapper: snake   # 表名映射规则（snake_case 到 CamelCase）
         column_mapper: snake  # 列名映射规则
     ```
     然后运行：
     ```
     reverse -f reverse.yml
     ```

3. **生成效果**：
   - 会自动生成结构体字段、类型映射（如int、varchar→string等）。
   - 默认添加`xorm` tag（如主键、唯一索引等）。
   - 支持多种数据库：MySQL、PostgreSQL、SQLite、MSSQL、Oracle等。

4. **注意事项**：
   - 生成的代码通常放在`models`目录下，你可以手动调整（如添加json tag等）。
   - xorm本身还支持**正向同步**（Sync2方法：从struct自动创建/更新表结构），但你问的是反向（表→struct），工具正是为此设计。
   - 官方文档：https://xorm.io/docs/ （搜索Reverse tool）。

如果你有具体的数据库类型或表名需求，可以提供更多细节，我可以帮你细化命令！