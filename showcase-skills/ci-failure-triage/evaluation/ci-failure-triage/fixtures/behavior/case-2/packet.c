uint32_t packet_size(const char *payload) {
  return (uint32_t)payload[0];
}

int main(void) {
  return (int)packet_size("x");
}
