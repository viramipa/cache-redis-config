import logging
import os
import redis
from typing import Optional

class RedisConfig:
    def __init__(self, host: str = 'localhost', port: int = 6379, db: int = 0):
        self.host = host
        self.port = port
        self.db = db
        self.redis_client = None

    def connect(self):
        try:
            self.redis_client = redis.Redis(host=self.host, port=self.port, db=self.db)
            self.redis_client.ping()
            logging.info(f"Connected to Redis at {self.host}:{self.port}")
        except redis.exceptions.ConnectionError as e:
            logging.error(f"Failed to connect to Redis: {e}")

    def get(self, key: str) -> Optional[str]:
        if self.redis_client:
            return self.redis_client.get(key)
        else:
            logging.error("Redis client is not connected")
            return None

    def set(self, key: str, value: str) -> bool:
        if self.redis_client:
            return self.redis_client.set(key, value)
        else:
            logging.error("Redis client is not connected")
            return False

    def delete(self, key: str) -> int:
        if self.redis_client:
            return self.redis_client.delete(key)
        else:
            logging.error("Redis client is not connected")
            return 0

def main():
    logging.basicConfig(level=logging.INFO)
    config = RedisConfig(host=os.environ.get('REDIS_HOST', 'localhost'), port=int(os.environ.get('REDIS_PORT', 6379)))
    config.connect()
    config.set('test_key', 'test_value')
    value = config.get('test_key')
    logging.info(f"Got value: {value}")
    config.delete('test_key')

if __name__ == '__main__':
    main()