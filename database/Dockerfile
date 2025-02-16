FROM postgres:17.2

# 環境変数を設定
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydatabase

# 初期データ投入用の SQL スクリプトをコピー
COPY init/* /docker-entrypoint-initdb.d/

# PostgreSQL を起動
EXPOSE 5432
CMD ["postgres"]

