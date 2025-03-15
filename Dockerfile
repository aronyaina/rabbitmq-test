# Use the official RabbitMQ image as the base
FROM rabbitmq:4.0

# Set a custom hostname (optional, can also be set in `docker run`)
ENV RABBITMQ_NODENAME rabbit@rabbit-1

# Enable RabbitMQ plugins (if needed)
RUN rabbitmq-plugins enable --offline rabbitmq_management

# Expose necessary ports
EXPOSE 15672 5672

