# Ground Truth

- Required outcome: every command consumer receives availability consistent with manifest-declared requirements and the active runtime.
- Observation point: `CommandPicker` displays unavailable commands.
- Current owners: the manifest owns requirements, `RuntimeCapabilities` owns runtime support, and `CommandRegistry` owns the shared command collection supplied to consumers.
- Smallest complete placement: make registry-level availability selection compose manifest requirements with runtime capabilities, inject runtime capabilities into the registry, and keep the picker as a consumer.
- Reconciliation: any registry consumers receive the same filtered set; avoid command-name switches or picker-local capability policy.
- Verification: multiple commands and capability sets, including no requirements and unavailable requirements, plus confirmation that the picker contains no command-specific rule.
